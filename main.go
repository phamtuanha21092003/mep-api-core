package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/phamtuanha21092003/mep-api-core/cmd/server"
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

	var runApp = &cobra.Command{
		Use:   "runserver [string to print]",
		Short: "Run gin app",
		Run: func(cmd *cobra.Command, args []string) {
			runGateway()
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
	}
	rootCmd.AddCommand(runApp)
	rootCmd.Execute()

}

// run gateway is run server
func runGateway() {
	config.LoadAllConfig()

	if database.SqlxConn == nil {
		database.NewDatabaseConn()
	}

	server := server.NewServer(database.SqlxConn)
	if err := server.RunServer(); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
