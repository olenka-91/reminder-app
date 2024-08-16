package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/olenka--91/reminder-app/internal/repository"
	"github.com/olenka--91/reminder-app/internal/service"
	"github.com/olenka--91/reminder-app/internal/transport/rest"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	if err := InitConfig(); err != nil {
		log.WithField("Err: ", err.Error()).Error("Couldn't read config")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: "qwerty",
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.WithField("Err:", err.Error()).Error("Couldn't create DB connection")
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handl := rest.NewHandler(serv)

	server := new(domain.Server)

	go func() {
		if err := server.Run(viper.GetString("port"), handl.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server %s", err.Error())
		}
	}()

	log.Println("Reminder Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Reminder Shutting Down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
