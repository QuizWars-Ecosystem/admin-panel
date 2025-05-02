package apis

import (
	"context"
	usersv1 "github.com/QuizWars-Ecosystem/admin-panel/gen/external/users/v1"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/clients"
	"net/mail"

	"go.uber.org/zap"
)

type Service struct {
	pool   *clients.GRPCClients
	logger *zap.Logger
}

func NewService(pool *clients.GRPCClients, logger *zap.Logger) *Service {
	return &Service{
		pool:   pool,
		logger: logger,
	}
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	res, err := s.pool.UsersClient.Login(ctx, &usersv1.LoginRequest{
		Identifier: &usersv1.LoginRequest_Email{
			Email: email,
		},
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return res.Token, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
