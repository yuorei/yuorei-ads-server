package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/yuovision-proto/go/video/video_grpc"
)

func (i *Infrastructure) UploadVideoForYuoVision(ctx context.Context, video *domain.UploadVideo, userID, uploadID, videoType string) error {
	stream, err := i.yuovision.VideoClient.UploadVideo(ctx)
	if err != nil {
		return err
	}

	meta := &video_grpc.UploadVideoInput_Meta{
		Meta: &video_grpc.VideoMeta{
			Id:                video.ID,
			Title:             video.Title,
			Description:       *video.Description,
			ThumbnailImageUrl: fmt.Sprintf("%s/%s/%s.webp", os.Getenv("AWS_S3_URL"), "thumbnail-image", video.ID),
			UserId:            userID,
			Tags:              video.Tags,
			Private:           video.IsPrivate,
			Adult:             video.IsAdult,
			ExternalCutout:    video.IsExternalCutout,
			IsAd:              video.IsAd,
		},
	}

	request := &video_grpc.UploadVideoInput{
		Value: meta,
	}

	err = stream.Send(request)
	if err != nil {
		return err
	}

	videoFile, err := os.Open("./uploads/" + uploadID + "_video." + videoType)
	if err != nil {
		return err
	}
	var data []byte

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(videoFile)
	if err != nil {
		return err
	}
	data = buf.Bytes()

	chunkSize := 3 * 1024 * 1024 // チャンクサイズ（3MB）
	for offset := 0; offset < len(data); offset += chunkSize {
		end := offset + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[offset:end]

		// Create a request containing the chunk of thumbnail data
		request := &video_grpc.UploadVideoInput{
			Value: &video_grpc.UploadVideoInput_Video{
				Video: chunk,
			},
		}
		// Send the chunk data
		err := stream.Send(request)
		if err != nil {
			return err
		}
	}

	// Receive response from the server
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}

func (i *Infrastructure) UploadThumbnailForYuoVision(ctx context.Context, thumbnail domain.ThumbnailImage, uploadID string) error {
	stream, err := i.yuovision.VideoClient.UploadThumbnail(ctx)
	if err != nil {
		return err
	}

	meta := &video_grpc.UploadThumbnailInput_Meta{
		Meta: &video_grpc.ThumbnailMeta{
			Id:          thumbnail.ID,
			ContentType: "image/" + thumbnail.ContentType,
		},
	}

	request := &video_grpc.UploadThumbnailInput{
		Value: meta,
	}

	err = stream.Send(request)
	if err != nil {
		return err
	}

	thumbnailImage, err := os.Open("./uploads/" + uploadID + "_image." + thumbnail.ContentType)
	if err != nil {
		return err
	}
	var data []byte

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(thumbnailImage)
	if err != nil {
		return err
	}
	data = buf.Bytes()

	chunkSize := 3 * 1024 * 1024 // チャンクサイズ（3MB）
	for offset := 0; offset < len(data); offset += chunkSize {
		end := offset + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[offset:end]

		// Create a request containing the chunk of thumbnail data
		request := &video_grpc.UploadThumbnailInput{
			Value: &video_grpc.UploadThumbnailInput_ThumbnailImage{
				ThumbnailImage: chunk,
			},
		}
		// Send the chunk data
		err := stream.Send(request)
		if err != nil {
			return err
		}
	}

	// Receive response from the server
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}
