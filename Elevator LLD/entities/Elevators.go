package entities

import "sync"

type Elevator struct {
	Buttons 	[]Button
	mu      	sync.RWMutex
	pos     	int
	target		int
	direction	bool
}