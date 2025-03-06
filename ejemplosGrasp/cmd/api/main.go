package main

import (
	"ejemplosGrasp/cmd/api/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	controller := controller.TestApi{}
	server := mux.NewRouter()
	server.HandleFunc("/controller/vehicles/{vehicle_id}/find", controller.FindControler).Methods("GET")
	server.HandleFunc("/polymorphism/vehicles/{model}/starts", controller.StartControler).Methods("GET")
	http.ListenAndServe(":8080", server)
}
