package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/StepanShevelev/l2/develop/timetable"
	"github.com/StepanShevelev/l2/develop/timetable/pkg/handler"
	"github.com/StepanShevelev/l2/develop/timetable/pkg/repository"
	"github.com/StepanShevelev/l2/develop/timetable/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("[cfg] error initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("[env] error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("[db] error initialization database: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(timetable.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.Logging(handlers.InitRouter())); err != nil {
			log.Fatalf("[srv] error running http server: %s", err.Error())
		}
	}()

	log.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)
	<-quit

	log.Print("Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("[srv] error shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Fatalf("[db] error connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
