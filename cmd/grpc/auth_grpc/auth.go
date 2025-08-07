package auth_grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"runtime/debug"
	"time"

	"github.com/phamtuanha21092003/mep-api-core/pkg/config"

	"google.golang.org/grpc"
)

func Run(ctx context.Context, name string) {
	serverAddr := fmt.Sprintf("%s:%d", config.AppGrpcCfg().Host, config.AppGrpcCfg().AuthPort)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic recovered: %v\n%s", r, debug.Stack())
			}
		}()

		log.Println("🚀 gRPC server started on " + serverAddr)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("✅ gRPC server stopped cleanly")

	case <-shutdownCtx.Done():
		log.Println("⏰ Timeout: forcing gRPC shutdown")
		grpcServer.Stop()
	}
}
