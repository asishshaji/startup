package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/asishshaji/startup/apps/auth/repository"
	"github.com/asishshaji/startup/apps/auth/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App creates the app
type App struct {
	httpServer *http.Server
	authUC     *usecase.AuthUseCase
}

// NewApp is the constructor
func NewApp() *App {
	db := initDB()

	userRepo := repository.NewUserRepository(db, "asd")
	return &App{
		authUC: usecase.NewAuthUseCase(*userRepo,
			"ASDS",
			[]byte("asd"),
			time.Hour*45,
		),
	}

}

// Run starts the server
func (*App) Run(port string) {

}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("adasd"))
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
