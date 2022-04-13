package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"microserviceSample/pkg/handlers"
	"microserviceSample/pkg/repositories"
	"microserviceSample/pkg/services"
)

var engine *gin.Engine

func run() {
	db := new(gorm.DB)
	repos := repositories.RepositoryInit(db)
	serv := services.ServiceInit(repos)
	handle := handlers.HandlerInit(serv)
	engine = handle.InitRoutes()

	err := engine.Run(viper.GetString("host") + ":" + viper.GetString("port"))
	if err != nil {
		log.Error().Msg("Cannot run web server")
		panic(err.Error())
	}

}

func main() {
	err := parseConfig()
	if err != nil {
		panic(err)
	}

	run()
}

func parseConfig() error {
	configPath := "../configuration"

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
