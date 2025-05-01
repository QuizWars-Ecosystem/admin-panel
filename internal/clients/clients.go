package clients

import (
	"context"
	usersv1 "github.com/QuizWars-Ecosystem/admin-panel/gen/external/users/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type UsersClient struct {
	usersv1.UsersAdminServiceClient
}

type GRPCClients struct {
	ctx     context.Context
	cancel  context.CancelFunc
	conn    *grpc.ClientConn
	address string
	opts    []grpc.DialOption
	mu      sync.RWMutex
	logger  *zap.Logger

	UsersClient *UsersClient
}

func NewGRPCClients(address string, logger *zap.Logger, opts ...grpc.DialOption) *GRPCClients {
	ctx, cancel := context.WithCancel(context.Background())

	c := &GRPCClients{
		ctx:         ctx,
		cancel:      cancel,
		address:     address,
		opts:        opts,
		logger:      logger,
		UsersClient: &UsersClient{},
	}

	go c.monitor()

	return c
}

func (c *GRPCClients) connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		_ = c.conn.Close()
	}

	c.logger.Info("Connecting to server", zap.String("address", c.address))

	conn, err := grpc.NewClient(c.address, c.opts...)
	if err != nil {
		return err
	}

	c.conn = conn

	c.UsersClient.UsersAdminServiceClient = usersv1.NewUsersAdminServiceClient(conn)

	c.logger.Info("Connected to server", zap.String("address", c.address))

	return nil
}

func (c *GRPCClients) monitor() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(time.Second * 5):
			c.mu.RLock()
			conn := c.conn
			c.mu.RUnlock()

			if conn == nil || conn.GetState().String() == "TRANSIENT_FAILURE" {
				c.logger.Debug("Reconnecting to server")
				_ = c.connect()
			}
		}
	}
}

func (c *GRPCClients) Close() {
	c.cancel()
}
