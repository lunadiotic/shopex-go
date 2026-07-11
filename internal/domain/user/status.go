package user

type Status string

const (
	StatusPending Status = "pending"
	StatusActive  Status = "active"
	StatusInactive Status = "inactive"
	StatusBlocked Status = "blocked"
)