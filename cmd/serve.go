package cmd

import (
	"github.com/kave08/url-shortener/config"
	handlers "github.com/kave08/url-shortener/handler"
	"github.com/kave08/url-shortener/repository"
	"github.com/kave08/url-shortener/routes"
	"github.com/kave08/url-shortener/service"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the url shortener service",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	dbs := config.LoadConfig(configPath)

	repos := repository.NewRepository(dbs.MysqlConnection, dbs.RedisConnection)
	lgcs := service.NewLogics(repos)
	handler := handlers.NewHandlers(lgcs)

	e := echo.New()
	e.HideBanner = false

	routes.InitializeGroup(e, handler)
	e.Logger.Fatal(e.Start(":8080"))
}

func init() {
	rootCMD.AddCommand(serveCmd)
}
