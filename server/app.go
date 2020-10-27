package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/asishshaji/startup/auth/repository"
	"github.com/asishshaji/startup/auth/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	httpServer *http.Server

	authUC usecase.AuthUseCase
}

func NewApp() *App {

	db := initDB()

	userRepo := repository.NewUserRepository(db,"collection")


	return &App{
		authUC: usecase.NewAuthUseCase(userRepo,
		"ASDS",
		"ASDASD",
		time.Hour* 45,
		),
	}
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("")
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("DB")
}
