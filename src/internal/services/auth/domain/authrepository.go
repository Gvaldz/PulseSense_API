package domain

import (
	user "pulse_sense/src/internal/users/domain"
)

type AuthRepository interface {
	FindUserByEmail(email string) (user.User, error)
	UpdateLastLogin(userID int32) error
}
