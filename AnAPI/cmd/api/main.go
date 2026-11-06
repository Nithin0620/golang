package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.setReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service on port 8000...")

	err := http.ListenAndServe("localhost:8000",r) 

	if err != nil {
		log.Error("Error starting server: %v", err)
	}
}