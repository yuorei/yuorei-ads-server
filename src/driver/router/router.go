package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/yuorei/yuorei-ads/gen/rpc/ads/v1/adsv1connect"
	"github.com/yuorei/yuorei-ads/gen/rpc/organization/v1/organizationv1connect"
	"github.com/yuorei/yuorei-ads/gen/rpc/user/v1/userv1connect"
	"github.com/yuorei/yuorei-ads/middleware"
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/adapter/presentation"
	"github.com/yuorei/yuorei-ads/src/adapter/presentation/rest"
	"github.com/yuorei/yuorei-ads/src/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewRouter() {
	infra := infrastructure.NewInfrastructure()
	repository := usecase.NewRepository(infra)

	// rest handler
	handler := rest.NewHandler(repository)

	// grpc handler
	ads := presentation.NewAdsServer(repository)
	user := presentation.NewUserServer(repository)
	organization := presentation.NewOrganizationServer(repository)

	firebaseApp, err := middleware.NewFirebaseApp()
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle(adsv1connect.NewAdManagementServiceHandler(ads))                   // TODO: interfaceを実装していく
	mux.Handle(userv1connect.NewUserServiceHandler(user))                         // TODO: interfaceを実装していく
	mux.Handle(organizationv1connect.NewOrganizationServiceHandler(organization)) // TODO: interfaceを実装していく

	mux.HandleFunc("/ad/video", handler.UploadAdVideoHandler)

	host := os.Getenv("IP") + ":" + os.Getenv("PORT")
	fmt.Println("Server is running on " + host)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*.yuorei.com", "http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
	})
	log.Fatalln(http.ListenAndServe(
		host,
		c.Handler(
			firebaseApp.AuthMiddleware(
				// Use h2c so we can serve HTTP/2 without TLS.
				// HTTP1.1リクエストはHTTP/2にアップグレードされる
				h2c.NewHandler(mux, &http2.Server{}),
			),
		),
	))
}
