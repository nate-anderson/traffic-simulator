package traffic

import (
	"fmt"
	"io"
)

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
		"[%d] Vehicle %d in '%s' :: lane '%s' (%s) => '%s' (%s)",
		m.Tick,
		m.Vehicle.Identify(),
		m.InJunction.Identifier,
		m.FromLane.Identify(),
		m.FromLane.Direction().String(),
		m.ToLane.Identify(),
		m.ToLane.Direction().String(),
	)
}

// SimulationReport is a sortable, filterable list of vehicle movements in a simulation
type SimulationReport []VehicleMovement

func (r SimulationReport) Write(writer io.WriteCloser) error {
	defer writer.Close()
	for _, line := range r {
		_, err := writer.Write([]byte(line.String() + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
