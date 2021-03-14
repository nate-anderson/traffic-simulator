package traffic

// Junction is the primary building block of traffic simulations.
// It represents a junction of multiple arms
type Junction struct {
	// Arms are the intersection's arms
	EnteringLanes []Lane
	// Select determines which lanes in each arm should proceed
	Proceed ProceedFunc
	// junction identifier
	Identifier string
}

// NewJunction instantiates a junction
func NewJunction(identifier string, proceed ProceedFunc) *Junction {
	return &Junction{
		Identifier: identifier,
		Proceed:    proceed,
	}
}

// AddEnteringLanes to the junction
func (j *Junction) AddEnteringLanes(lanes ...Lane) {
	for _, lane := range lanes {
		j.EnteringLanes = append(j.EnteringLanes, lane)
	}
}

// ProceedFunc is a function type for determining which lanes should allow traffic to proceed
type ProceedFunc func(lane Lane) (shouldProceed bool)
