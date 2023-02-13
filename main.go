package main

import (
	"fmt"
	"handson/internal/config"
	"handson/internal/handler"
	"handson/internal/logging"
	"handson/internal/repository"
	"handson/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conf := config.Config()

	db := sqlx.MustConnect("mysql", conf.DBSrc())
	defer db.Close()

	logger := logging.Logger()

	r := mux.NewRouter()

	userUsecase := usecase.NewUser(repository.NewUser(), db)
	r.HandleFunc("/users", handler.PostUser(userUsecase, logger)).Methods("POST")

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	loggerRouter := logging.Middleware(logger)(r)

	http.ListenAndServe(fmt.Sprintf(":%s", config.Config().Port), loggerRouter)
}
