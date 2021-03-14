package main

import "github.com/nate-anderson/traffic-simulator"

/*
Simple intersection with one one-way input and three one-way outputs
Input from the south, outputs to the west, north, east
*/

// SimpleIntersection demonstrates a simple intersection simulation
func SimpleIntersection() {
	// create lanes
	input := traffic.NewFIFOLane(traffic.DirectionN, "northbound input")
	outputNorth := traffic.NewFIFOLane(traffic.DirectionN, "northbound output")
	outputWest := traffic.NewFIFOLane(traffic.DirectionW, "westbound output")
	outputEast := traffic.NewFIFOLane(traffic.DirectionE, "eastbound output")

	input.NextLanes = []traffic.Lane{
		outputNorth,
		outputWest,
		outputEast,
	}

	incomingVehicles := traffic.MakeNVehicles(100)
	input.AddInitialVehicles(incomingVehicles...)

	junction := traffic.Junction{
		Lanes: []traffic.Lane{
			input, outputNorth, outputWest, outputEast,
		},
		Identifier: "Simple Intersection",
		AllowsLane: func(l traffic.Lane) bool {
			// north input lane is the only input lane, always allow it to proceed
			return l.Direction() == traffic.DirectionN
		},
	}

	sim := traffic.NewSimulation("Simple intersection", 5, 5)
	sim.AddJunction(&junction)
	sim.Run()
}
