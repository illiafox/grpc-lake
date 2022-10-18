package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"server/internal/adapters/api"
	"server/internal/controller/grpc"
	httpserver "server/internal/controller/http"
)

func (app *App) Listen(item api.ItemUsecase) {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// // gRPC
	grpcServer := grpc.NewServer(app.Logger, item)
	go func() {
		defer cancel()
		app.Logger.Info("gRPC server started", zap.Int("port", app.Config.GRPC.Port))

		err := grpcServer.Listen(app.Config.GRPC.Port)
		if err != nil {
			app.Logger.Error("grpc server", zap.Error(err))
		}
	}()

	// HTTP
	httpServer := httpserver.NewServer(app.Logger, "0.0.0.0", app.Config.HTTP.Port, item)
	go func() {
		defer cancel()
		app.Logger.Info("HTTP server started", zap.Int("port", app.Config.HTTP.Port))

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
			app.Logger.Error("http server", zap.Error(err))
		}
	}()

	<-ctx.Done()
	_, _ = os.Stdout.WriteString("\n")

	// // Graceful shutdown

	// HTTP
	app.Logger.Info("Shutting down HTTP server")
	err := httpServer.Shutdown(context.TODO())
	if err != nil {
		app.Logger.Error("shutdown http server", zap.Error(err))
	}

	// gRPC
	app.Logger.Info("Shutting down gRPC server")
	grpcServer.GracefulStop()
}
