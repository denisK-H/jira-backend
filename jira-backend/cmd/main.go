package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"

	"github.com/JingolBong/jira-connector/internal/db"
	"github.com/JingolBong/jira-connector/jira-backend/internal/handler"
	"github.com/JingolBong/jira-connector/jira-backend/internal/repository"
	"github.com/JingolBong/jira-connector/jira-backend/internal/service"
)
// Инициализация сервера, DI зависимостей и регистрация роутов (TODO: конфиг, таймауты, логирование)
func main() {

	dsn := "postgres://pguser:pgpwd@localhost:5432/testdb?sslmode=disable" // TODO: переиспользовать из dbConnection.go

	writeDB, _ := sql.Open("postgres", dsn)
	readDB, _ := sql.Open("postgres", dsn)

	storage := db.NewStorage(writeDB, readDB)

	repo := repository.NewProjectRepository(storage, writeDB)

	projectService := service.NewProjectService(repo)
	issueService := service.NewIssueService(repo)
	graphService := service.NewGraphService(repo)

	projectHandler := handler.NewProjectHandler(projectService)
	issueHandler := handler.NewIssueHandler(issueService)
	graphHandler := handler.NewGraphHandler(graphService)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/projects", projectHandler.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/projects/{id:[0-9]+}", projectHandler.Delete).Methods("DELETE")

	r.HandleFunc("/api/v1/issues", issueHandler.GetByProject).Methods("GET")

	r.HandleFunc("/api/v1/graph/make/{task:[0-9]+}", graphHandler.Make).Methods("POST")
	r.HandleFunc("/api/v1/graph/get/{task:[0-9]+}", graphHandler.Get).Methods("GET")
	r.HandleFunc("/api/v1/isAnalyzed", graphHandler.IsAnalyzed).Methods("GET")

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}