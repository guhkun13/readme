package entity

type Attendant struct {
	userID string
	name   string
}

type Invitation struct {
	attendant Attendant
	event     Event
	isConfirm bool
}
