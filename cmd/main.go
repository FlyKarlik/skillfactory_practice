package main

import (
	"log"
	"os"
	"skillfactory_project/pkg/handler"
	"skillfactory_project/pkg/repository"
	service "skillfactory_project/pkg/sercvice"
	"skillfactory_project/pkg/server"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db, err := repository.NewPostgresDb(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	mongoClient, err := repository.NewMongoDb(os.Getenv("MONGO_URL"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db, mongoClient)
	//В сервисе меняем базу данных(в этом примере mongodb)
	service := &service.Sercive{
		PostsService: service.NewMongoService(repo.PostsRepositoryMongo),
	}
	handler := handler.NewHandler(service)
	server := server.NewServer()
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}

}
