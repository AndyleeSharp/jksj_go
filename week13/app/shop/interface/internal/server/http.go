package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "week13.myapp/api/shop/interface/v1"
	"week13.myapp/app/shop/interface/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/shop.interface.v1.ShopInterface/Login"] = struct{}{}
	whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(s *service.ShopInterface) *http.Server {
	var opts = []http.ServerOption{

		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}

	opts = append(opts, http.Address("0.0.0.0:8000"))
	srv := http.NewServer(opts...)
	// h := openapiv2.NewHandler()
	// srv.HandlePrefix("/q/", h)
	v1.RegisterShopInterfaceHTTPServer(srv, s)
	return srv
}
