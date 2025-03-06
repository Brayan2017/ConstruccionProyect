package vehicleServices

import (
	"ejemplosGrasp/internal/domain"
	"ejemplosGrasp/internal/interfaces"
)

type VehicleService struct{}

func (u VehicleService) GetModelService(id int) domain.Vehicle {
	return domain.NewVehicle(id, "model", "color")
}

func (u VehicleService) StartVehicleService(vehicle interfaces.Start) string {
	return vehicle.StartVehicle()
}
