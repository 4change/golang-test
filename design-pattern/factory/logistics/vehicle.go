package main

type Vehicle struct {
	RoadLogistics
}

func NewVehicle() ILogistics {
	return &Vehicle{ 
		RoadLogistics: RoadLogistics{ 
			name: "vehicle",
		},
	}
}