package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/phamtuanha21092003/mep-api-core/pkg/config"
	"github.com/phamtuanha21092003/mep-api-core/pkg/middleware"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
)

// Inital config application
func init() {
	godotenv.Load()
	env := os.Getenv("APP_ENV")
	config.AppConfig = config.NewConfiguration(env)
}

func main() {
	config.LoadAllConfig()

	if database.SqlxConn == nil {
		database.NewDatabaseConn()
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.New()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	middleware.GinMiddleware(router, config.AppConfig)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

}
