package main 	

import (
	"fmt" 
	"net/http" 
	// package for handling HTTP 
	"github.com/go-chi/chi"
	"github.com/KaiHodges/go-API/internal/handlers"
	// below sets log as the alias 
	log "github.com/sirupsen/logrus" 
)


func main(){
	
	log.SetReportCaller(true) 
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service... ")
	
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
 
