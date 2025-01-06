package client

import (
	"log"
	"os"

	"github.com/yuorei/yuorei-ads/yuovision-proto/go/video/video_grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientYuoVision struct {
	VideoClient   video_grpc.VideoServiceClient
	UserClient    video_grpc.UserServiceClient
	CommentClient video_grpc.CommentServiceClient
}

func NewClientYuoVision() *ClientYuoVision {
	client := &ClientYuoVision{}
	client.NewConnectYuoVision()
	return client
}

func (c *ClientYuoVision) NewConnectYuoVision() {
	address := os.Getenv("YUOVISION_ADDRESS")
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal("Connection failed. err: ", err)
		return
	}

	c.VideoClient = video_grpc.NewVideoServiceClient(conn)
	c.UserClient = video_grpc.NewUserServiceClient(conn)
	c.CommentClient = video_grpc.NewCommentServiceClient(conn)
}
