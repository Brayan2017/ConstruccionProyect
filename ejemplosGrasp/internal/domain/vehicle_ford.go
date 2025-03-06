package domain

type VehicleFord struct {
	Vehicle
}

func (v VehicleFord) StartVehicle() string {
	return "Starting vehicle Ford with button"
}
