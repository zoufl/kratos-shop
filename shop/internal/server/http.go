package server

import (
	"context"
	v1 "shop/api/shop/v1"
	"shop/internal/conf"
	"shop/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, shop *service.ShopService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(jwt.Server(func(token *jwt2.Token) (any, error) {
				return []byte(ac.JwtKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
			).Match(NewWhiteListMatcher()).Build(),
			logging.Server(logger),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterShopHTTPServer(srv, shop)
	return srv
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/shop.shop.v1.Shop/Captcha"] = struct{}{}
	whiteList["/shop.shop.v1.Shop/Login"] = struct{}{}
	whiteList["/shop.shop.v1.Shop/Register"] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}

		return true
	}
}
