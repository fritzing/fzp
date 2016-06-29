package fzp

type Buses []Bus

func NewBuses() *Buses {
	b := Buses{}
	b = make([]Bus, 0)
	return &b
}
