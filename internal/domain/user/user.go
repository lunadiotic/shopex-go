package user

import "time"

type User struct {
	ID uint
	FullName string
	Email string
	Password string
	
	Status Status

	Roles []Role

	CreatedAt time.Time	
	UpdatedAt time.Time

}

func (u User) Hasrole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}