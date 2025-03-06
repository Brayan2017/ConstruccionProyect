package domain

type Vehicle struct {
	Id    int
	Model string
	Color string
}

func NewVehicle(id int, model string, color string) Vehicle {
	return Vehicle{
		Id:    id,
		Model: model,
		Color: color,
	}
}
