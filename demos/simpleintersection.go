package main

import (
	"os"

	"github.com/nate-anderson/traffic-simulator"
)

/*
Simple intersection with one one-way input and three one-way outputs
Input from the south, outputs to the west, north, east
*/

// simpleIntersection demonstrates a simple intersection simulation
func simpleIntersection() {
	// create lanes
	input := traffic.NewFIFOLane("northbound input", traffic.DirectionN)
	outputNorth := traffic.NewFIFOLane("northbound output", traffic.DirectionN)
	outputWest := traffic.NewFIFOLane("westbound output", traffic.DirectionW)
	outputEast := traffic.NewFIFOLane("eastbound output", traffic.DirectionE)

	input.NextLanes = []traffic.Lane{
		outputNorth,
		outputWest,
		outputEast,
	}

	incomingVehicles := traffic.MakeNVehicles(100)
	input.AddInitialVehicles(incomingVehicles...)

	junction := traffic.Junction{
		EnteringLanes: []traffic.Lane{
			input, outputNorth, outputWest, outputEast,
		},
		Identifier: "Simple Intersection",
		Proceed: func(l traffic.Lane) bool {
			// north input lane is the only input lane, always allow it to proceed
			return l.Direction() == traffic.DirectionN
		},
	}

	sim := traffic.NewSimulation("Simple intersection", 5, 5)
	sim.AddJunctions(&junction)
	report := sim.Run()
	report.Write(os.Stdout)
}
