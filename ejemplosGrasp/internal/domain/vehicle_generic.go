package domain

type VehicleGeneric struct {
	Vehicle
}

func (v VehicleGeneric) StartVehicle() string {
	return "Starting vehicle Generic with key"
}
