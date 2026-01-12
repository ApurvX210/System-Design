package entities

import "sync"

type Elevator struct {
	id				int
	Buttons 		[]Pannel
	mu      		sync.RWMutex
	pos     		int
	targetQueue		[]int
	direction		bool
}

func (el *Elevator) TargetAdd(target int){
	el.mu.Lock()
	defer el.mu.Unlock()

	el.targetQueue = append(el.targetQueue, target)
}

func (el *Elevator) GetELevatorId() int{
	el.mu.RLocker()
	defer el.mu.RUnlock()

	return el.id
}

func (el *Elevator) GetPosAndTarget() (int,bool,[]int){
	el.mu.RLocker()
	defer el.mu.RUnlock()

	return el.pos,el.direction,el.targetQueue
}
