package main

import (
	"flag"
	"fmt"
	"net/http"

	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/backend-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Add("Access-Control-Allow-Headers", "Authorization")
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
	// 	switch e := err.(type) {
	// 	case *response.CodeError:
	// 		return http.StatusOK, e.Data()
	// 	default:
	// 		return http.StatusInternalServerError, nil
	// 	}
	// })

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
