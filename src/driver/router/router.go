package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yuorei/yuorei-ads/gen/rpc/ads/v1/adsv1connect"
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/adapter/presentation"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewRouter() {
	infra := infrastructure.NewInfrastructure()
	ads := presentation.NewAdsServer(infra)
	mux := http.NewServeMux()
	path, handler := adsv1connect.NewAdManagementServiceHandler(ads) // TODO: interfaceを実装していく
	mux.Handle(path, handler)

	host := os.Getenv("IP") + ":" + os.Getenv("PORT")
	fmt.Println("Server is running on " + host)
	http.ListenAndServe(
		host,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
