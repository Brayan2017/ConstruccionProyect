package controller

import (
	"ejemplosGrasp/internal/domain"
	"ejemplosGrasp/internal/interfaces"
	vehicleServices "ejemplosGrasp/internal/services/vehicle_services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TestApi struct{}

func (c TestApi) FindControler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleId, err := strconv.Atoi(params["vehicle_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Vehicle invalid id error "))
		return
	}

	vehicleData := vehicleServices.VehicleService{}.GetModelService(vehicleId)
	vehicle, _ := json.Marshal(vehicleData)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("vehicle valid id: " + string(vehicle)))
}

func (c TestApi) StartControler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	model := params["model"]

	var instance interfaces.Start

	if model == "generic" {
		instance = domain.VehicleGeneric{}
	} else if model == "ford" {
		instance = domain.VehicleFord{}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Vehicle invalid model error "))
		return
	}

	vehicleData := vehicleServices.VehicleService{}.StartVehicleService(instance)

	w.Write([]byte((vehicleData)))
}
