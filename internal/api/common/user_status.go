package common

type UserStatus int

const (
	UserStatusActive   UserStatus = 0
	UserStatusInactive UserStatus = 1
	UserStatusBanned   UserStatus = 2
	UserStatusDeleted  UserStatus = 3
	UserStatusUnknown  UserStatus = -1
)
