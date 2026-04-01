package main

import (
	"fmt"
	"log"
	"net/http"

	"courses/internal/config"
	"courses/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	env := config.LoadEnv()

	db, err := config.InitDB(env.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Falha crítica ao conectar no banco: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	_ = userRepo

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API rodando com variáveis de ambiente configuradas!"))
	})

	port := fmt.Sprintf(":%s", env.APIPort)
	fmt.Printf("🚀 Servidor iniciado na porta %s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("❌ Erro ao iniciar servidor: %v", err)
	}
}
