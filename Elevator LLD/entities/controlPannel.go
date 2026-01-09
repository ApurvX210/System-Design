package entities

type ButtonType string

const (
	FLOOR     ButtonType = "FLOOR"
	EMERGENCY ButtonType = "EMERGENCY"
	CONTACT   ButtonType = "CONTACT"
)

type Button struct {
	Type ButtonType
}

type Pannel struct {
	OnClick (Button)
}

type ExternalPannel struct {
	floorId  	int
	buttonList  []Button
	signalChan	chan Signal
}

func (pannel ExternalPannel) OnClick(b Button) {

}

type InternalPannel struct {
	buttonList []Button
	signalChan	chan Signal
}

func (pannel InternalPannel) OnClick(b Button) {

}