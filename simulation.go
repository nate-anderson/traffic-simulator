package traffic

import "fmt"

// Simulation runs a traffic simulation with one or more junctions
type Simulation struct {
	Name            string
	Junctions       []*Junction
	VehiclesPerTick int
	Ticks           int
}

// NewSimulation creates a new simulation
func NewSimulation(name string, vehiclesPerTick int, ticks int) *Simulation {
	return &Simulation{
		Name:            name,
		VehiclesPerTick: vehiclesPerTick,
		Ticks:           ticks,
	}
}

// AddJunction to a simulation
func (s *Simulation) AddJunction(j *Junction) {
	s.Junctions = append(s.Junctions, j)
}

// Run runs the simulation as configured
func (s Simulation) Run() {
	for i := 0; i < s.Ticks; i++ {
		for _, junction := range s.Junctions {
			fmt.Printf("Junction %s\n", junction.Identifier)
			for _, lane := range junction.EnteringLanes {
				// if the junction allows the lane to pass traffic, and the lane has a destination lane
				// to pass traffic to, handle updates to the lane
				if junction.AllowsLane(lane) && lane.HasDestination() {
					for i := 0; i < s.VehiclesPerTick; i++ {
						destLanes := lane.DestinationLanes()
						crossingVehicle, exists := lane.GetDeparture()
						if exists {
							crossingVehicle.DoVisit(junction)
							// determine which lane the vehicle selects, and pass it along
							targetLane := crossingVehicle.SelectLane(destLanes)
							targetLane.GiveArrival(crossingVehicle)
							fmt.Printf("Vehicle %d in junction %s: moves to lane %s (direction %s)\n", crossingVehicle.Identify(), junction.Identifier, targetLane.Identify(), targetLane.Direction().String())
						}
					}
				}
			}
		}
	}
}
