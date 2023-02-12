package main

import (
	"context"
	config2 "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"imageUpload/internal/config"
	"imageUpload/internal/repository"
	service2 "imageUpload/internal/service"
	http2 "imageUpload/internal/transport/http"
	"imageUpload/pkg/database"
	"imageUpload/pkg/hash"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg := config.New()

	log.Printf("config: %+v\n", cfg)

	db, err := database.NewPostgresDB(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	defer db.Close()

	s3cfg, s3err := config2.LoadDefaultConfig(context.TODO())
	if s3err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := s3.NewFromConfig(s3cfg)

	hasher := hash.New(cfg.Phrase.Salt)

	userRepo := repository.New(db)
	userService := service2.NewUserRepo(userRepo, hasher, []byte(cfg.Phrase.Secret))
	handler := http2.NewHandler(userService, client)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
