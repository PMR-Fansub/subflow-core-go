package constants

type UserStatus int

const (
	UserStatusUnknown  UserStatus = 0
	UserStatusActive   UserStatus = 1
	UserStatusInactive UserStatus = 2
	UserStatusBanned   UserStatus = 3
	UserStatusDeleted  UserStatus = 4
)
