# Traffic Simulator

This package provides a framework for simulating traffic systems. It attempts to be general-purpose and unopinionated, but is very new and may have significant shortcomings. Simulations are non-realtime. This is just something I'm experimenting with and interested in learning about,

## How to use this

There are several examples included in the `/demos` directory.

A simulation consists of one or more junctions, and runs for a predefined number of ticks.

```go
// a simulation that will run for 50 ticks
sim := traffic.NewSimulation("Example Simulation", 1, 50)
```

A junction is a meeting of one or more roads - it can represent an at-grade or grade-separated interchange such as an intersection, a roundabout, and others.

```go
// create a junction
jxn := traffic.NewJunction("Test Intersection")
```

### Junctions

A junction consists of two or more lanes that feed traffic into it. Any lane feeding traffic into a junction must have one or more destination lanes which receive its traffic. The destination lanes should not also belong to the junction unless they also feed traffic into it (which I suspect is a rare situation).

#### Lanes

`Lane` is an interface which can be implemented many different ways. The `traffic` package exports a default implementation of `Lane`, called `FIFOLane`, which implements a simple first-in, first-out queue lane. `Lane` is currently a rather large interface, and `FIFOLane` will fulfill most use cases that I can foresee.

#### Directions

Lanes must indicate their direction, which can be one of the standard cardinal directions exported by the package (`DirectionN`, `DirectionE`, `DirectionNW`, etc.), or any `int` value cast to a `Direction` type if more directions are needed.

```go
// create lanes
n3right := traffic.NewFIFOLane("Northbound 3rd St. Right Lane", traffic.DirectionN)
w12right := traffic.NewFIFOLane("Westbound 12th St. Right Lane", traffic.DirectionE)

// add the lanes as inputs to a junction
jxn.AddEnteringLanes(n3right, w12right)
```

Every lane feeding into a junction must point to the lanes into which its traffic must flow. The vehicles passing through the first lane will decide which destination lane they enter on the other side of the junction.

```go
/// w12right is a turn lane feeding into n3right
w12right.TurnLane = true
w12right.AddDestination(n3right)
```

#### Proceed Function

Junctions require a function to control the flow of traffic through itself (of type `traffic.ProceedFunc`). This function is called on each lane feeding into the intersection and returns a boolean indicating if traffic is permitted to flow. The library does not come with a default implementation of this function because it is entirely model-specific.

Suggestions for implementing this function:

- Use a struct to keep track of intersection-specific variables and use the struct method as the Junction decision function
- Simple junctinos with only north-south and east-west lanes may alternate between even- and odd-directioned lanes. An example is included below.

```go
// example ProceedFunc
func simpleProceed(lane Lane) bool {
    return lane.Direction() % 2 == 0
}
```

### Vehicles

Vehicles are mobile actors which move from lane to lane through junctions in the simulation. `Vehicle` is an interface that may be implemented to create detailed movement reports, but the package includes the `DefaultVehicle` implementation which keeps track of the junctions it visits.

At the beginning of the simulation, there must be at least one vehicle placed in a connected lane in order for the simulation to run.

```go
// populate a lane with 10 vehicles
vehicles := traffic.MakeNVehicles(10)
n3right.AddInitialVehicles(vehicles...)
```
