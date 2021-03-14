package traffic

import (
	"container/list"
	"fmt"
)

// Lane represents a lane on a road, which can produce and (optionally) receive traffic
type Lane interface {
	// Produce a vehicle
	GetDeparture() (Vehicle, bool)
	// Pass a vehicle into this lane
	GiveArrival(v Vehicle)
	// GetQueueSize returns the number of vehicles waiting in this lane
	GetQueueSize() int
	// IsTurnLane indicates if this is a turn lane
	IsTurnLane() bool
	// Direction gets the lane's direction indicator
	Direction() Direction
	// Next points to the lane this lane outputs to. must be a part of the same junction
	DestinationLanes() []Lane
	// Add a possible lane for this lane to feed into
	AddDestination(lanes ...Lane)
	// Does the lane point to a receiving lane?
	AddInitialVehicles(vehicles ...Vehicle)
	HasDestination() bool
	Identify() string
}

// FIFOLane is a standard, first-in first-out lane
type FIFOLane struct {
	TurnLane  bool
	NextLanes []Lane
	direction Direction
	queue     *list.List
	name      string
}

// NewFIFOLane instantiates a new FIFO lane
func NewFIFOLane(name string, direction Direction) *FIFOLane {
	return &FIFOLane{
		direction: direction,
		queue:     list.New(),
		name:      name,
	}
}

// AddInitialVehicles to the lane
func (l *FIFOLane) AddInitialVehicles(vehicles ...Vehicle) {
	for _, v := range vehicles {
		l.queue.PushBack(v)
	}
}

// GetDeparture returns the next vehicle in the lane queue, if any
func (l FIFOLane) GetDeparture() (v Vehicle, hasVehicle bool) {
	el := l.queue.Front()
	if el == nil {
		hasVehicle = false
		return
	}
	l.queue.Remove(el)
	v, ok := el.Value.(Vehicle)
	if !ok {
		err := fmt.Errorf("value stored in FIFOLane queue was not a Vehicle: type %T found", v)
		panic(err)
	}
	hasVehicle = true
	return

}

// GiveArrival passes an entering vehicle into the lane
func (l *FIFOLane) GiveArrival(v Vehicle) {
	l.queue.PushBack(v)
}

// GetQueueSize of the lane's queue
func (l *FIFOLane) GetQueueSize() int {
	return l.queue.Len()
}

// IsTurnLane indicates if the lane is a turn lane
func (l *FIFOLane) IsTurnLane() bool {
	return l.TurnLane
}

// DestinationLanes returns the lanes which this lane feeds to
func (l FIFOLane) DestinationLanes() []Lane {
	return l.NextLanes
}

// Direction the lane permits traffic
func (l FIFOLane) Direction() Direction {
	return l.direction
}

// HasDestination - does the lane have a destination for exiting traffic?
func (l FIFOLane) HasDestination() bool {
	return len(l.NextLanes) > 0
}

// AddDestination adds possible outbound lanes following this one
func (l *FIFOLane) AddDestination(lanes ...Lane) {
	for _, lane := range lanes {
		l.NextLanes = append(l.NextLanes, lane)
	}
}

// Identify the lane by name
func (l FIFOLane) Identify() string {
	return l.name
}

// Direction is a direction a lane may be running
type Direction int

const (
	// DirectionN is the package direction constant for North
	DirectionN Direction = iota
	// DirectionE is the package direction constant for East
	DirectionE
	// DirectionS is the package direction constant for South
	DirectionS
	// DirectionW is the package direction constant for West
	DirectionW
	// DirectionNE is the package direction constant for Northeast
	DirectionNE
	// DirectionSE is the package direction constant for Southeast
	DirectionSE
	// DirectionSW is the package direction constant for Southwest
	DirectionSW
	// DirectionNW is the package direction constant for Northwest
	DirectionNW
)

func (d Direction) String() string {
	switch d {
	case DirectionN:
		return "North"
	case DirectionE:
		return "East"
	case DirectionS:
		return "South"
	case DirectionW:
		return "West"
	case DirectionNE:
		return "Northeast"
	case DirectionSE:
		return "Southeast"
	case DirectionSW:
		return "Southwest"
	case DirectionNW:
		return "Northwest"
	default:
		return fmt.Sprintf("Custom direction %d", d)
	}
}
