package main

import (
    "github.com/your_project/internal/repository/postgres"
    "github.com/your_project/internal/service"
    "github.com/your_project/internal/transport/http"
    "github.com/your_project/pkg/db"
    "log"
)

func main() {
    connectionString := "postgresql://user:password@localhost:5432/yourdb?sslmode=disable"
    db := db.NewPostgresDB(connectionString)

    repo := postgres.NewEstimationRepository(db)
    svc := service.NewEstimationService(repo)
    router := http.NewRouter(svc)

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("could not run server: %v", err)
    }
}
