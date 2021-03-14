package traffic

import "fmt"

// VehicleMovement is a record of a vehicle's movement in one tick
type VehicleMovement struct {
	Tick       int
	Vehicle    Vehicle
	InJunction *Junction
	FromLane   Lane
	ToLane     Lane
}

func (m VehicleMovement) String() string {
	return fmt.Sprintf(
		"[%d] Vehicle %d/Junction '%s' :: lane '%s' (%s) => '%s' (%s)",
		m.Tick,
		m.Vehicle.Identify(),
		m.InJunction.Identifier,
		m.FromLane.Identify(),
		m.FromLane.Direction().String(),
		m.FromLane.Identify(),
		m.FromLane.Direction().String(),
	)
}

// SimulationReport is a sortable, filterable list of vehicle movements in a simulation
type SimulationReport []VehicleMovement
