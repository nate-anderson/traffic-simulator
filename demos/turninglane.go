package main

import "github.com/nate-anderson/traffic-simulator"

/*
Turning lane leading to an eastbound turn. Demonstrates use of `traffic.Junction` for added lanes.

                |
        ________/
_______/___________

*/

// turningLane demonstrates a straight single-lane road leading to an intersection,
// adding a turn lane leading up to the turn
func turningLane() {
	eastboundPreTurn := traffic.NewFIFOLane("eastbound, before turn lane", traffic.DirectionE)
	eastboundNorthTurnLane := traffic.NewFIFOLane("eastbound, north turn lane", traffic.DirectionE)
	eastboundPostTurnLane := traffic.NewFIFOLane("eastbound, after turn lane", traffic.DirectionE)

	// eastbound splits into straight and turn lane
	eastboundPreTurn.AddDestination(eastboundNorthTurnLane, eastboundPostTurnLane)

	// create turn lane junction, always allowing traffic through
	turnLaneJxn := traffic.NewJunction("eastbound north turn lane", func(l traffic.Lane) bool {
		return true
	})
	turnLaneJxn.AddEnteringLanes(eastboundPreTurn)

	northbound := traffic.NewFIFOLane("nortbound", traffic.DirectionN)
	eastboundPostIntersection := traffic.NewFIFOLane("eastbound, after intersection", traffic.DirectionE)

	eastboundNorthTurnLane.AddDestination(northbound)
	eastboundPostTurnLane.AddDestination(eastboundPostIntersection)

	// create simple intersection with proceedFunc that always allows eastbound lanes to proceed
	jxn := traffic.NewJunction("simple junction", func(lane traffic.Lane) bool {
		if lane.Direction() == traffic.DirectionE {
			return true
		}
		return false
	})

	jxn.AddEnteringLanes(eastboundNorthTurnLane, eastboundPostTurnLane)

	// add vehicles to initial eastbound lane
	eastboundPreTurn.AddInitialVehicles(
		traffic.MakeNVehicles(20)...,
	)

	// create simulation
	sim := traffic.NewSimulation("Turn lane", 2, 10)
	sim.AddJunctions(turnLaneJxn, jxn)
	sim.Run()
}
