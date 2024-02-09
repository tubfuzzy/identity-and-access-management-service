package service

import "authentication-and-authorization-service/internal/domain"

type authenticationService struct {
	domain.AuthenticationService
}

func NewAuthenticationService(authenticationRepository domain.AuthenticationRepository) domain.AuthenticationService {
	return &authenticationService{
		AuthenticationService: authenticationRepository,
	}
}
