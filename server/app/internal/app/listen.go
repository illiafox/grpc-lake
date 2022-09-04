package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"server/app/internal/controller/grpc"
	http_server "server/app/internal/controller/http"
	"server/app/pkg/log"
)

func (app *App) Listen() {
	defer app.closers.Close(app.logger)

	item, err := app.ItemService()
	if err != nil {
		app.logger.Error("create item service", log.Error(err))
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	// // gRPC
	grpcServer := grpc.NewServer(app.logger, item)
	go func() {
		defer cancel()
		app.logger.Info("gRPC server started", log.Int("port", app.cfg.GRPC.Port))

		err := grpcServer.Listen(app.cfg.GRPC.Port)
		if err != nil {
			app.logger.Error("grpc server", log.Error(err))
		}
	}()

	// HTTP
	httpServer := http_server.NewServer("0.0.0.0", app.cfg.HTTP.Port)
	go func() {
		defer cancel()
		app.logger.Info("HTTP server started", log.Int("port", app.cfg.HTTP.Port))

		err := httpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			app.logger.Error("http server", log.Error(err))
			return
		}
	}()

	<-ctx.Done()
	_, _ = os.Stdout.WriteString("\n")

	// // Graceful shutdown

	// HTTP
	app.logger.Info("Shutting down HTTP server")
	err = httpServer.Shutdown(context.TODO())
	if err != nil {
		app.logger.Error("shutdown http server", log.Error(err))
	}

	// gRPC
	app.logger.Info("Shutting down gRPC server")
	grpcServer.GracefulStop()
}
