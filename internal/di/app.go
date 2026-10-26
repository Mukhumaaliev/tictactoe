package di

import (
	"context"
	"net"
	"net/http"
	"os"
	"tictactoe/internal/datasource"
	"tictactoe/internal/domain"
	"tictactoe/internal/web"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		datasource.NewDatabase,
		datasource.NewConfig,
		datasource.NewGameRepository,
		domain.NewService,
		web.NewGameHandler,
		web.NewRouter,
		NewHTTPServer,
	),
	fx.Invoke(RegisterHooks),
)

func NewHTTPServer(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr: func() string {
			if p := os.Getenv("PORT"); p != "" {
				return ":" + p
			}
			return ":8081"
		}(),
		Handler: mux,
	}
}

func RegisterHooks(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go func() {
				_ = server.Serve(ln)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
