package main

type Train struct {
	RoadLogistics
}

func NewTrain() ILogistics {
	t := new(Train)
	t.RoadLogistics = RoadLogistics{
		name: "train",
	}
	return t
}