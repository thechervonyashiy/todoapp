package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	todoapp "github.com/thechervonyashiy/todoapp"
	"github.com/thechervonyashiy/todoapp/pkg/handlers"
	"github.com/thechervonyashiy/todoapp/pkg/lib/logger"
	"github.com/thechervonyashiy/todoapp/pkg/repository"
	"github.com/thechervonyashiy/todoapp/pkg/routes"
	"github.com/thechervonyashiy/todoapp/pkg/services"
)

func main() {
	log := logger.SetupLogger()

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error to loading env veriables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("faild to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)
	router := routes.InitRoutes(handler)

	srv := new(todoapp.Server)
	if err := srv.Run(viper.GetString("port"), router); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
