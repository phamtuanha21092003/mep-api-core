package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/cmd/server/dependencies"
	"github.com/phamtuanha21092003/mep-api-core/cmd/server/router"
	"github.com/phamtuanha21092003/mep-api-core/pkg/config"
	"github.com/phamtuanha21092003/mep-api-core/pkg/middleware"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
	"github.com/phamtuanha21092003/mep-api-core/platform/logger"
)

type Server struct {
	gin          *gin.Engine
	sqlx         *database.SqlxDatabase
	logger       *logger.Logger
	controller   *dependencies.Controllers
	repositories dependencies.Repositories
	services     dependencies.Services
}

func NewServer(db *database.SqlxDatabase) *Server {
	logger.SetUpLogger()
	logger := logger.GetLogger()

	engine := gin.New()
	middleware.GinMiddleware(engine, config.AppConfig)

	return &Server{gin: engine, sqlx: db, logger: logger}
}

func (server *Server) RunServer() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server.setupRoutes()

	server.logger.Printf("Server running.....")
	serverAddr := fmt.Sprintf("%s:%d", config.AppCfg().Host, config.AppCfg().Port)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: server.gin,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("listen: %s\n", err)
				panic(err)
			}
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func (server *Server) setupRoutes() {
	router.GeneralRouter(server.gin)
	router.UserRouter(server.gin)
}
