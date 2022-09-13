package main

import "errors"

func getLogistics(means string) (i ILogistics, err error){
	if  means == "vehicle" {
		return NewVehicle(), nil
	}else if means == "train"{
		return NewTrain(), nil
	}
	return nil, errors.New("unknown means of transportation")
}