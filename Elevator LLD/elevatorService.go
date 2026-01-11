package main

import (
	"context"

	"./entities"
)

type ElevatorService struct {
	signalChan	chan entities.Signal
	// pannel   	[]*entities.ExternalPannel
	elevator 	[]*entities.Elevator
}

func (es *ElevatorService) HandleSignal(signal entities.Signal) error{
	return nil
}

func (es *ElevatorService) ListeningSignal(ctx context.Context){
	select{
	case signal := <- es.signalChan:
		go es.HandleSignal(signal)
	case <- ctx.Done():
		return
	}
}
