package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/martonkaufmann/erp/handler/customer"
	"github.com/martonkaufmann/erp/middleware"
	"github.com/martonkaufmann/erp/provider"
)

func main() {
	r := http.NewServeMux()

	customer.RegisterRoutes(r)

	s := http.Server{
		Addr:              ":8080",
		Handler:           middleware.RequestLog(middleware.JSON(r)),
		WriteTimeout:      time.Second * 5,
		ReadTimeout:       time.Second * 15,
		IdleTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Millisecond * 500,
		BaseContext: func(listener net.Listener) context.Context {
			ctx := context.Background()
			ctx = provider.WithLog(ctx)
			ctx = provider.WithDatabase(ctx)
			ctx = provider.WithValidate(ctx)

			return ctx
		},
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
