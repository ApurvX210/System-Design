package entities

import "sync"

type Elevator struct {
	Buttons 	[]Pannel
	mu      	sync.RWMutex
	pos     	int
	target		int
	direction	bool
}