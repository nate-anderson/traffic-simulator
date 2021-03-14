package traffic

import "math/rand"

// Vehicle is a vehicle in traffic
type Vehicle interface {
	// @TODO this should log a visit to each intersection
	DoVisit(j *Junction)
	// identify the vehicle
	Identify() int
	// select a lane when presented with a choice
	SelectLane(lanes []Lane) Lane
}

// DefaultVehicle is the package's default Vehicle implementation
type DefaultVehicle struct {
	ID              int
	junctionHistory []string
}

// DoVisit records a vehicle's visit to a junction
func (v *DefaultVehicle) DoVisit(j *Junction) {
	v.junctionHistory = append(v.junctionHistory, j.Identifier)
}

// Identify identifies the vehicle
func (v DefaultVehicle) Identify() int {
	return v.ID
}

// SelectLane randomly selects one of the provided lanes
func (v DefaultVehicle) SelectLane(lanes []Lane) Lane {
	i := rand.Intn(len(lanes))
	return lanes[i]
}

var currentID = 1

// MakeNVehicles makes a set of unique default vehicles
func MakeNVehicles(n int) []Vehicle {
	all := make([]Vehicle, 0, n)

	for i := 0; i < n; i++ {
		v := DefaultVehicle{
			ID:              currentID,
			junctionHistory: []string{},
		}
		all = append(all, &v)
		currentID++
	}
	return all
}
