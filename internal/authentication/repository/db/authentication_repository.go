package db

import "authentication-and-authorization-service/internal/domain"

type authenticationRepository struct {
}

func NewAuthenticationRepository() domain.AuthenticationRepository {
	return &authenticationRepository{}
}
