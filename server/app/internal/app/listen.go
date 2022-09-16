package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"server/app/internal/adapters/api"
	"server/app/internal/controller/grpc"
	httpserver "server/app/internal/controller/http"
	"server/app/pkg/log"
)

func (app *App) Listen(item api.ItemUsecase) {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// // gRPC
	grpcServer := grpc.NewServer(app.Logger, item)
	go func() {
		defer cancel()
		app.Logger.Info("gRPC server started", log.Int("port", app.Config.GRPC.Port))

		err := grpcServer.Listen(app.Config.GRPC.Port)
		if err != nil {
			app.Logger.Error("grpc server", log.Error(err))
		}
	}()

	// HTTP
	httpServer := httpserver.NewServer("0.0.0.0", app.Config.HTTP.Port)
	go func() {
		defer cancel()
		app.Logger.Info("HTTP server started", log.Int("port", app.Config.HTTP.Port))

		var err error

		if app.Config.Flags.HTTPS {
			err = httpServer.ListenAndServeTLS(
				app.Config.HTTP.HTTPS.CertFile,
				app.Config.HTTP.HTTPS.KeyFile,
			)
		} else {
			err = httpServer.ListenAndServe()
		}

		if err != http.ErrServerClosed {
			app.Logger.Error("http server", log.Error(err))
			return
		}
	}()

	<-ctx.Done()
	_, _ = os.Stdout.WriteString("\n")

	// // Graceful shutdown

	// HTTP
	app.Logger.Info("Shutting down HTTP server")
	err := httpServer.Shutdown(context.TODO())
	if err != nil {
		app.Logger.Error("shutdown http server", log.Error(err))
	}

	// gRPC
	app.Logger.Info("Shutting down gRPC server")
	grpcServer.GracefulStop()
}
