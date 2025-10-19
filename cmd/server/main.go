package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/configs"
	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/webserver/handlers"
	"github.com/pablorodrigovieira/go-expert/challenges/go-multithreading/internal/webserver/services"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	services.Init(config)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/cep/{cep}", handlers.GetCep)

	http.ListenAndServe(":8000", router)
}
