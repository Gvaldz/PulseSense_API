package application

import (
	"errors"
	"pulse_sense/src/core"
	auth "pulse_sense/src/internal/services/auth/domain"
	user "pulse_sense/src/internal/users/domain"
)

type Login struct {
	authRepo     auth.AuthRepository
	userRepo     user.UserRepository
	tokenService auth.TokenService
	hasher       core.PasswordHasher
}

func NewLogin(
	authRepo auth.AuthRepository,
	userRepo user.UserRepository,
	tokenService auth.TokenService,
	hasher core.PasswordHasher,
) *Login {
	return &Login{
		authRepo:     authRepo,
		userRepo:     userRepo,
		tokenService: tokenService,
		hasher:       hasher,
	}
}

func (uc *Login) Execute(credentials user.User) (auth.Token, int32, error) {
	user, err := uc.authRepo.FindUserByEmail(credentials.Correo)
	if err != nil {
		return auth.Token{}, 0, errors.New("datos incorrectos")
	}

	if err := uc.hasher.Compare(user.Contrasena, credentials.Contrasena); err != nil {
		return auth.Token{}, 0, errors.New("datos incorrectos")
	}

	token, err := uc.tokenService.GenerateToken(user.IdUsuario, user.Correo, user.Tipo)
	if err != nil {
		return auth.Token{}, 0, errors.New("fallo en generar token")
	}

	go func() {
		_ = uc.authRepo.UpdateLastLogin(user.IdUsuario)
	}()

	return token, user.Tipo, nil
}