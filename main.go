package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	authGrpc "github.com/phamtuanha21092003/mep-api-core/cmd/grpc/auth_grpc"
	"github.com/phamtuanha21092003/mep-api-core/cmd/server"
	syncPermission "github.com/phamtuanha21092003/mep-api-core/cmd/sync_permission"
	"github.com/phamtuanha21092003/mep-api-core/pkg/config"
	"github.com/phamtuanha21092003/mep-api-core/platform/database"
)

// Inital config application
func init() {
	godotenv.Load()
	env := os.Getenv("APP_ENV")
	config.AppConfig = config.NewConfiguration(env)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var runAppCmd = &cobra.Command{
		Use:   "runserver",
		Short: "Run gin app",
		Run: func(cmd *cobra.Command, args []string) {
			runGateway()
		},
	}

	var runGrpcCmd = &cobra.Command{
		Use:   "grpc [name ...]",
		Short: "Run grpc app",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadAllConfigGrpc()

			if database.SqlxConn == nil {
				database.NewDatabaseConn()
			}

			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer stop()

			if len(args) == 0 {
				log.Fatal("Please provide a gRPC service name or 'all'")
			}

			services := map[string]func(context.Context, string){
				"auth": authGrpc.Run,
			}

			if len(args) == 0 {
				log.Fatal("Please provide a gRPC service name or 'all'")
			}

			// all only for development
			if len(args) == 1 && config.AppGrpcCfg().Environment == "dev" && args[0] == "all" {
				var wg sync.WaitGroup
				for name, run := range services {
					wg.Add(1)
					go func(name string, runFunc func(context.Context, string)) {
						defer wg.Done()
						log.Printf("🚀 Starting gRPC service: %s", name)
						runFunc(ctx, name)
					}(name, run)
				}
				wg.Wait()
				return
			}

			for _, name := range args {
				run, exists := services[name]
				if !exists {
					log.Fatalf("❌ Unknown service: %s", name)
				}
				log.Printf("🚀 Starting gRPC service: %s", name)
				run(ctx, name)
			}
		},
	}

	var syncPermCmd = &cobra.Command{
		Use:   "sync-permissions",
		Short: "Sync Gin routes as permissions into DB",
		Run: func(cmd *cobra.Command, args []string) {
			syncPermissions()
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
	}
	rootCmd.AddCommand(runAppCmd, runGrpcCmd, syncPermCmd)
	rootCmd.Execute()

}

// run gateway is run server
func runGateway() {
	config.LoadAllConfigServer()

	if database.SqlxConn == nil {
		database.NewDatabaseConn()
	}

	server := server.NewServer(database.SqlxConn)
	if err := server.RunServer(); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func syncPermissions() {
	config.LoadAllConfigServer()

	if database.SqlxConn == nil {
		database.NewDatabaseConn()
	}

	syncPermission.SyncPermission(database.SqlxConn)
}
