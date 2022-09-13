package main

import (
	"fmt"
)

func main() {
    vehicle, _ := getLogistics("vehicle")
    train, _ := getLogistics("train")

    printDetails(vehicle)
    printDetails(train)
}

func printDetails(g ILogistics) {
    fmt.Printf("Vehicle: %s\n", g.getName())
}