package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase"
)

type Handler struct {
	usecase *usecase.UseCase
}

func NewHandler(repository *usecase.Repository) *Handler {
	return &Handler{
		usecase: usecase.NewUseCase(repository),
	}
}

func (h *Handler) UploadAdVideoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	userID, ok := ctx.Value("uid").(string)
	if !ok || userID == "" {
		http.Error(w, "Unable to get userID", http.StatusForbidden)
		return
	}

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		http.Error(w, "Unable to create uploads directory", http.StatusInternalServerError)
		return
	}

	// クエリパラメータからファイル名とチャンク番号を取得
	videoType := r.FormValue("videoType")
	videoTypeSplited := strings.Split(videoType, "/")
	if len(videoTypeSplited) != 2 {
		http.Error(w, "Invalid video type", http.StatusBadRequest)
		return
	}
	videoType = videoTypeSplited[1]
	imageType := r.FormValue("imageType")
	imageType = strings.Split(imageType, "/")[1]

	chunkNumberStr := r.FormValue("chunkNumber")
	totalChunksStr := r.FormValue("totalChunks")

	uploadID := r.FormValue("uploadID")

	// チャンク番号を整数に変換
	chunkNumber, err := strconv.Atoi(chunkNumberStr)
	if err != nil {
		http.Error(w, "Invalid chunk number", http.StatusBadRequest)
		return
	}

	// 画像ファイルを取得
	imageFile, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the image file", http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	// 動画ファイルを取得
	videoFile, _, err := r.FormFile("video")
	if err != nil {
		http.Error(w, "Error retrieving the video file", http.StatusBadRequest)
		return
	}
	defer videoFile.Close()

	// 画像チャンクの一時保存ファイル名を生成
	imageTempFileName := fmt.Sprintf("./uploads/%s_image.part%d", uploadID, chunkNumber)
	imageDst, err := os.Create(imageTempFileName)
	if err != nil {
		http.Error(w, "Unable to create the image chunk file", http.StatusInternalServerError)
		return
	}
	defer imageDst.Close()

	// 動画チャンクの一時保存ファイル名を生成
	videoTempFileName := fmt.Sprintf("./uploads/%s_video.part%d", uploadID, chunkNumber)
	videoDst, err := os.Create(videoTempFileName)
	if err != nil {
		http.Error(w, "Unable to create the video chunk file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(videoTempFileName)
	defer videoDst.Close()

	// 画像ファイルの内容をコピー
	if _, err := io.Copy(imageDst, imageFile); err != nil {
		http.Error(w, "Unable to save the image chunk", http.StatusInternalServerError)
		return
	}

	// 動画ファイルの内容をコピー
	if _, err := io.Copy(videoDst, videoFile); err != nil {
		http.Error(w, "Unable to save the video chunk", http.StatusInternalServerError)
		return
	}
	// チャンクがすべてアップロードされたか確認
	totalChunks, err := strconv.Atoi(totalChunksStr)
	if err != nil {
		http.Error(w, "Invalid total chunks", http.StatusBadRequest)
		return
	}

	if chunkNumber == totalChunks-1 {
		// すべてのチャンクがアップロードされた後、ファイルを再構成する
		finalImageFileName := "./uploads/" + uploadID + "_image" + "." + imageType
		finalImageFile, err := os.Create(finalImageFileName)
		if err != nil {
			http.Error(w, "Unable to create final image file", http.StatusInternalServerError)
			return
		}
		defer os.Remove(finalImageFileName)
		defer finalImageFile.Close()

		finalVideoFileName := "./uploads/" + uploadID + "_video" + "." + videoType
		finalVideoFile, err := os.Create(finalVideoFileName)
		if err != nil {
			http.Error(w, "Unable to create final video file", http.StatusInternalServerError)
			return
		}
		defer finalVideoFile.Close()

		for i := 0; i < totalChunks; i++ {
			imageTempFileName := fmt.Sprintf("./uploads/%s_image.part%d", uploadID, i)
			imageTempFile, err := os.Open(imageTempFileName)
			if err != nil {
				http.Error(w, "Unable to open image chunk file", http.StatusInternalServerError)
				return
			}

			if _, err := io.Copy(finalImageFile, imageTempFile); err != nil {
				http.Error(w, "Unable to append image chunk to final file", http.StatusInternalServerError)
				imageTempFile.Close()
				return
			}
			imageTempFile.Close()

			// 画像チャンクファイルを削除
			err = os.Remove(imageTempFileName)
			if err != nil {
				http.Error(w, "Unable to delete image chunk file", http.StatusInternalServerError)
			}
		}

		// 各動画チャンクを結合
		for i := 0; i < totalChunks; i++ {
			videoTempFileName := fmt.Sprintf("./uploads/%s_video.part%d", uploadID, i)
			videoTempFile, err := os.Open(videoTempFileName)
			if err != nil {
				http.Error(w, "Unable to open video chunk file", http.StatusInternalServerError)
				return
			}

			if _, err := io.Copy(finalVideoFile, videoTempFile); err != nil {
				http.Error(w, "Unable to append video chunk to final file", http.StatusInternalServerError)
				videoTempFile.Close()
				return
			}
			videoTempFile.Close()

			// 動画チャンクファイルを削除
			os.Remove(videoTempFileName)
		}

		// フォームからAdVideoMetaの情報を取得
		title := r.FormValue("title")
		description := r.FormValue("description")
		thumbnailImageUrl := r.FormValue("thumbnail_image_url")
		tagString := r.FormValue("tags")
		tags := strings.Split(tagString, ",")
		campaignID := r.FormValue("campaign_id")
		adLink := r.FormValue("ad_link")

		adID := domain.NewAdID()
		ad := domain.NewAd(adID, campaignID, "video", false, false, adLink, tags, time.Now(), time.Now(), nil)

		adVideo := domain.NewAdVideo(adID, title, description, "", thumbnailImageUrl, time.Now(), time.Now(), nil)

		adResult, err := h.usecase.AdsInputPort.CreateAdVideo(ctx, ad, adVideo, userID, uploadID, videoType, imageType)
		if err != nil {
			http.Error(w, "Unable to process video meta", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(struct {
			AdID string `json:"adID,omitempty"`
		}{adResult.AdID})
		if err != nil {
			http.Error(w, "Unable to encode response", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
