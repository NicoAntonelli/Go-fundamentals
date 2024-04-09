package main

import (
	"fmt"
	"net/http"
	"simple-api/internals/handlers"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Constants
const SERVER = "localhost:8080"
const VERSION = "v1.1.0"

func main() {
	log.SetReportCaller(true)
	router := mux.NewRouter()
	handlers.Handler(router)

	fmt.Printf("Starting GO API Service...\n")

	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/                                                    
	`)

	fmt.Printf("Simple-API project %v running on: %v\n", VERSION, SERVER)

	err := http.ListenAndServe(SERVER, router)
	if err != nil {
		log.Error(err)
	}
}
