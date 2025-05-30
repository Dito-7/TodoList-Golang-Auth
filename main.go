package main

import (
	"TodoList-Golang-Auth/delivery"
	"TodoList-Golang-Auth/repository/mongodb"
	"TodoList-Golang-Auth/routes"
	"TodoList-Golang-Auth/usecase"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		slog.Warn(".env file not found, using default environment variables")
	} else {
		slog.Info(".env file loaded successfully")
	}
}

func mongoConnection() (*mongo.Client, error) {
	mongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		return nil, fmt.Errorf("MONGO_URI tidak ditemukan di .env")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke MongoDB: %v", err)
	}

	dbName := os.Getenv("MONGO_DBNAME")
	if err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, fmt.Errorf("gagal ping MongoDB: %v", err)
	}

	slog.Info("Berhasil terhubung ke MongoDB")
	return client, nil
}

func main() {
	dbClient, err := mongoConnection()
	if err != nil {
		slog.Error("Error koneksi MongoDB", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			slog.Error("Error saat menutup koneksi MongoDB", "error", err)
		}
	}()

	database := dbClient.Database(os.Getenv("MONGO_DBNAME"))

	blacklistRepo := mongodb.NewBlacklistRepository(database)

	if err := blacklistRepo.EnsureTTLIndex(context.TODO()); err != nil {
		slog.Error("Gagal membuat TTL index untuk blacklist tokens", "error", err)
		os.Exit(1)
	}

	userRepo := mongodb.NewUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := delivery.NewUserHandler(userUsecase, blacklistRepo)

	todoRepo := mongodb.NewTodoRepository(database)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := delivery.NewTodoHandler(todoUsecase)

	r := chi.NewRouter()
	routes.SetupUserRoutes(r, userHandler, todoHandler, blacklistRepo)

	port := ":4444"
	slog.Info("Server started", "port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		slog.Error("Gagal menjalankan server", "error", err)
	}
}
