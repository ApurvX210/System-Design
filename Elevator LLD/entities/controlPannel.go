package entities

type ButtonType string

const (
	FLOOR     ButtonType = "FLOOR"
	EMERGENCY ButtonType = "EMERGENCY"
	CONTACT   ButtonType = "CONTACT"
)

type Button struct {
	Type 	ButtonType
	Value	int
}

type Pannel struct {
	OnClick (Button)
}

type ExternalPannel struct {
	floorId  	int
	buttonList  []Button
	signalChan	chan Signal
}

func (pannel *ExternalPannel) OnClick(b Button) {
	signal := Signal{
		floor: pannel.floorId,
		button: b,
	}
	pannel.signalChan <- signal
}

type InternalPannel struct {
	floorId  	int
	buttonList []Button
	signalChan	chan Signal
}

func (pannel *InternalPannel) OnClick(b Button) {
	signal := Signal{
		floor: pannel.floorId,
		button: b,
	}
	pannel.signalChan <- signal
}