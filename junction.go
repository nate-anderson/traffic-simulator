package traffic

// Junction is the primary building block of traffic simulations.
// It represents a junction of multiple arms
type Junction struct {
	// Arms are the intersection's arms
	Lanes []Lane
	// Select determines which lanes in each arm should proceed
	AllowsLane ProceedFunc
	// junction identifier
	Identifier string
}

// ProceedFunc is a function type for determining which lanes should allow traffic to proceed
type ProceedFunc func(lane Lane) (shouldProceed bool)
