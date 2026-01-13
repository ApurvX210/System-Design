package main

import (
	"context"
	"./entities"
)

type ElevatorService struct {
	signalChan	chan entities.Signal
	// pannel   	[]*entities.ExternalPannel
	elevators 	[]*entities.Elevator
}

func (es *ElevatorService) HandleSignal(signal entities.Signal) (entities.Elevator,error){
	sameDir := []*entities.Elevator{}
	stillElev := []*entities.Elevator{}
	oppostiteDir := []*entities.Elevator{}
	crntFloor := signal.GetFloorId()
	targetFloor := signal.GetButton().Value
	for _,elv := range es.elevators{
		
	}
}

func (es *ElevatorService) ListeningSignal(ctx context.Context){
	select{
	case signal := <- es.signalChan:
		go es.HandleSignal(signal)
	case <- ctx.Done():
		return
	}
}
