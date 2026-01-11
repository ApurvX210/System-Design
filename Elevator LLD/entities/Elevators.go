package entities

import "sync"

type Elevator struct {
	Buttons 		[]Pannel
	mu      		sync.RWMutex
	pos     		int
	targetQueue		[]int
	direction		bool
}