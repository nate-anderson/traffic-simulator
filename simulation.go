package traffic

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

// AddJunctions to a simulation
func (s *Simulation) AddJunctions(junctions ...*Junction) {
	for _, j := range junctions {
		s.Junctions = append(s.Junctions, j)
	}
}

// Run runs the simulation as configured
func (s Simulation) Run() SimulationReport {
	report := SimulationReport{}
	for tick := 0; tick < s.Ticks; tick++ {
		for _, junction := range s.Junctions {
			for _, lane := range junction.EnteringLanes {
				// if the junction allows the lane to pass traffic, and the lane has a destination lane
				// to pass traffic to, handle updates to the lane
				if junction.Proceed(lane) && lane.HasDestination() {
					for i := 0; i < s.VehiclesPerTick; i++ {
						destLanes := lane.DestinationLanes()
						crossingVehicle, exists := lane.GetDeparture()
						if exists {
							crossingVehicle.DoVisit(junction)
							// determine which lane the vehicle selects, and pass it along
							targetLane := crossingVehicle.SelectLane(destLanes)
							targetLane.GiveArrival(crossingVehicle)
							reportEntry := VehicleMovement{
								Tick:       tick,
								Vehicle:    crossingVehicle,
								InJunction: junction,
								FromLane:   lane,
								ToLane:     targetLane,
							}
							report = append(report, reportEntry)
						}
					}
				}
			}
		}
	}
	return report
}
