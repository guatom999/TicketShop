package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guatom999/TicketShop/config"
	"github.com/guatom999/TicketShop/modules/middlewares/middlewaresUseCases"
	"github.com/guatom999/TicketShop/modules/middlewares/middlwareHandlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type (
	Server interface {
		Start(pctx context.Context)
	}

	server struct {
		app        *echo.Echo
		db         *gorm.DB
		cfg        *config.Config
		middleware middlwareHandlers.MiddlwareHandlerService
	}
)

func NewMiddleware(cfg *config.Config) middlwareHandlers.MiddlwareHandlerService {
	usecase := middlewaresUseCases.NewMiddlewareUsecase()
	handler := middlwareHandlers.NewMiddlewareHandler(cfg, usecase)

	return handler
}

func NewEchoServer(db *gorm.DB, cfg *config.Config) Server {
	return &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: NewMiddleware(cfg),
	}
}

func (s *server) gracefulShutdown(pctx context.Context, close <-chan os.Signal) {

	<-close

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown")
		panic(err)
	}

	log.Println("Shutting down Server....")

}

func (s *server) Start(pctx context.Context) {

	s.app.Use(middleware.Logger())

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(pctx, close)

	log.Println("Starting server...")

	s.userModule()

	if err := s.app.Start(fmt.Sprintf(":%d", s.cfg.App.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to shutdown:%v", err)

	}
}
