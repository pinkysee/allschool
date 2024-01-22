package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/PINKYSEE/schoolserver/internal/app/apiserver"
	"github.com/PINKYSEE/schoolserver/internal/app/handler"
	"github.com/PINKYSEE/schoolserver/internal/app/repository"
	"github.com/PINKYSEE/schoolserver/internal/app/services"
	"github.com/PINKYSEE/schoolserver/internal/logger"
	"github.com/PINKYSEE/schoolserver/internal/storage/database"
)

var (
	configPath          string
	database_configPath string
	logger_configPath   string
	hand                *handler.Handling
)

func init() {
	flag.StringVar(&configPath, "config-path", "../../configs/apiserver.toml", "path to cfg file")
	flag.StringVar(&database_configPath, "database_config-path", "../../configs/database.toml", "path to db_cfg file")
	flag.StringVar(&logger_configPath, "logger_config-path", "../../configs/logger.toml", "path to logger_cfg file")
}
func main() {
	flag.Parse()
	loggerinit()
	db_Configinit()

	configinit()
}

func configinit() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config, hand)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

func db_Configinit() {
	config := database.NewConfig()
	_, err := toml.DecodeFile(database_configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Start(config)
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(db)
	serv := services.NewServices(repos)
	hand = handler.Handlinginit(serv)
}

func loggerinit() {
	config := logger.NewConfig()
	_, err := toml.DecodeFile(logger_configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	if err := logger.Start(config); err != nil {
		log.Fatal(err)
	}
}
