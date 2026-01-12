package entities

type Signal struct{
	floor	int
	button  Button
}

func (s *Signal) GetFloorId() int{
	return s.floor
}

func (s *Signal) GetButton() Button{
	return s.button
}