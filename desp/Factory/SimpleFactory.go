package Factory

import "learndesign/desp/models"

const (
	FrontUser = iota
	AdminUser
)

type UserType int

func CreateUser(t UserType) models.UserCreateFunc {

	switch t {
	case FrontUser:
		return models.NewUser()
	case AdminUser:
		return models.NewAdmin()
	default:
		return models.NewUser()

	}

}
