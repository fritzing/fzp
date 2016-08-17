package fzp

// Buses store all Bus objects as an array.
type Buses []Bus

// NewBuses return a new Buses object.
func NewBuses() *Buses {
	b := Buses{}
	b = make([]Bus, 0)
	return &b
}
