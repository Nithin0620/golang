package main

import (
	"fmt"
	"github.com/Nithin0620/golang/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service on port 8000...")

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Errorf("Error starting server: %v", err)
	}
}
