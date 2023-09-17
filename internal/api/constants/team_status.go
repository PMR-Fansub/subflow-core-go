package constants

type TeamStatus int

const (
	TeamStatusUnknown TeamStatus = 0
	TeamStatusOpen    TeamStatus = 1
	TeamStatusClosed  TeamStatus = 2
	TeamStatusBanned  TeamStatus = 3
	TeamStatusDeleted TeamStatus = 4
)
