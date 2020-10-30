package server

import (
	"context"
	"log"
	"time"

	"github.com/asishshaji/startup/apps/auth/delivery"
	"github.com/asishshaji/startup/apps/auth/repository"
	"github.com/asishshaji/startup/apps/auth/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App creates the app
type App struct {
	httpRouter delivery.Router
	authUC     *usecase.AuthUseCase
	port       string
}

// NewApp is the constructor
func NewApp(router *delivery.Router, port string) *App {
	db := initDB()

	userRepo := repository.NewUserRepository(db, "asd")
	userUseCase := usecase.NewAuthUseCase(*userRepo,
		"ASDS",
		[]byte("asd"),
		time.Hour*45,
	)
	return &App{
		httpRouter: *router,
		authUC:     userUseCase,
		port:       port,
	}

}

// Run starts the server
func (app *App) Run() {

	// app.httpRouter.POST("/signup")
	// app.httpRouter.POST("/signin")
	app.httpRouter.SERVE(app.port)
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
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

	log.Println("Connected to MongoDB ")

	return client.Database("DB")
}