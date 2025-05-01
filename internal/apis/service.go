package apis

import (
	"context"
	"github.com/QuizWars-Ecosystem/admin-panel/internal/clients"

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

func (s *Service) Login(ctx context.Context, identifier, password string) error {

	return nil
}
