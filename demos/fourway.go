package main

import (
	"math/rand"

	"github.com/nate-anderson/traffic-simulator"
)

/*
Simulate an intersection of two roads, each with two lanes passing each direction.
*/

// a traffic signal to alternate between north-south and east-west traffic
type alternatingTrafficSignal struct {
	northSouth bool
}

func (s *alternatingTrafficSignal) Proceed(lane traffic.Lane) bool {
	if s.northSouth {
		s.northSouth = !s.northSouth
		return lane.Direction() == traffic.DirectionN || lane.Direction() == traffic.DirectionS
	}

	s.northSouth = !s.northSouth
	return lane.Direction() == traffic.DirectionE || lane.Direction() == traffic.DirectionW

}

// fourWayIntersection demos a two-lane, four-way intersection with each lane permitting
// either straight or turning traffic
func fourWayIntersection() {
	// create lanes
	eastboundInN := traffic.NewFIFOLane("eastbound north", traffic.DirectionE)
	eastboundInS := traffic.NewFIFOLane("eastbound south", traffic.DirectionE)
	eastboundOutN := traffic.NewFIFOLane("eastbound north", traffic.DirectionE)
	eastboundOutS := traffic.NewFIFOLane("eastbound south", traffic.DirectionE)

	westboundInN := traffic.NewFIFOLane("westbound north", traffic.DirectionW)
	westboundInS := traffic.NewFIFOLane("westbound south", traffic.DirectionW)
	westboundOutN := traffic.NewFIFOLane("westbound north", traffic.DirectionW)
	westboundOutS := traffic.NewFIFOLane("westbound south", traffic.DirectionW)

	northboundInE := traffic.NewFIFOLane("northbound east", traffic.DirectionN)
	northboundInW := traffic.NewFIFOLane("northbound west", traffic.DirectionN)
	northboundOutE := traffic.NewFIFOLane("northbound east", traffic.DirectionN)
	northboundOutW := traffic.NewFIFOLane("northbound west", traffic.DirectionN)

	southboundInE := traffic.NewFIFOLane("southbound east", traffic.DirectionS)
	southboundInW := traffic.NewFIFOLane("southbound west", traffic.DirectionS)
	southboundOutE := traffic.NewFIFOLane("southbound east", traffic.DirectionS)
	southboundOutW := traffic.NewFIFOLane("southbound west", traffic.DirectionS)

	incomingLanes := []traffic.Lane{
		eastboundInN,
		eastboundInS,
		westboundInN,
		westboundInS,
		northboundInE,
		northboundInW,
		southboundInE,
		southboundInW,
	}

	// link inbound lanes to outbound lanes
	eastboundInN.AddDestination(northboundOutW, eastboundOutN)
	eastboundInS.AddDestination(eastboundOutS, southboundOutW)

	westboundInN.AddDestination(westboundOutN, northboundOutE)
	westboundInS.AddDestination(westboundOutS, southboundOutE)

	northboundInE.AddDestination(northboundOutE, eastboundOutS)
	northboundInW.AddDestination(northboundOutW, westboundOutS)

	southboundInE.AddDestination(southboundOutE, eastboundOutN)
	southboundInW.AddDestination(southboundOutW, westboundOutN)

	// create traffic signal for the junction
	signal := alternatingTrafficSignal{false}

	// create junction
	intersection := traffic.NewJunction("four-way", signal.Proceed)

	intersection.AddEnteringLanes(incomingLanes...)

	// add some vehicles to initial simulation state in incoming lanes
	maxPerLane := 15
	for _, incoming := range incomingLanes {
		n := rand.Intn(maxPerLane + 1)
		vehicles := traffic.MakeNVehicles(n)
		incoming.AddInitialVehicles(vehicles...)
	}

	// create simulation, 10 ticks, 3 vehicles per tick
	sim := traffic.NewSimulation("Four-way Intersection Simulation", 3, 10)
	sim.AddJunctions(intersection)
	sim.Run()

}
