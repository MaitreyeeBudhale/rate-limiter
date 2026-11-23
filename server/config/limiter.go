package config

import "time"

type Limiter interface {
	Allow() bool //funciton with return type boolean
}

// Go doesn't require you to explicitly write something like:
// implements Limiter
// Go interfaces are satisfied implicitly.
// The compiler determines whether a type satisfies an interface based on its methods.

type Config struct {
	TotalRequests int
	Workers       int
	Capacity      int
	RefillRate    int
	RefillPeriod  time.Duration
}

type Stats struct {
	Accepted int64
	Rejected int64
}

// Limiter → the behaviour every algorithm must provide.
// Config → the settings that control the simulation and algorithm.
// Stats → the results of processing requests.
