package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/services/auth/application"
	"pulse_sense/src/internal/services/auth/infrastructure/controllers"
)

type AuthDependencies struct {
	DB       *sql.DB
	Hasher   *core.BcryptHasher
	UserRepo *core.UserRepository
}

func NewAuthDependencies(db *sql.DB, hasher *core.BcryptHasher, userRepo *core.UserRepository) *AuthDependencies {
	return &AuthDependencies{
		DB:       db,
		Hasher:   hasher,
		UserRepo: userRepo,
	}
}

func (d *AuthDependencies) GetRoutes() *AuthRoutes {
	authRepo := core.NewAuthRepository(d.DB)
	tokenService := core.NewJWTService()

	loginUC := application.NewLogin(
		authRepo,
		d.UserRepo,
		tokenService,
		d.Hasher,
	)

	loginController := controllers.NewLoginController(loginUC)

	return NewAuthRoutes(loginController)
}
