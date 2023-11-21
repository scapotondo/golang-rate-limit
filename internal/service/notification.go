package service

import (
	"context"
	"fmt"
	"golang-rate-limit/internal/domain"
	"golang-rate-limit/internal/logs"
	"golang-rate-limit/internal/resources"
	"sync"
)

type Notification interface {
	SendEmail(c context.Context, request resources.NotificationRequest) error
}

type notification struct {
	logger logs.Logger
	mu     sync.Mutex
	users  map[string]*domain.NotificationRateLimit
}

func (ns *notification) SendEmail(ctx context.Context, request resources.NotificationRequest) error {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	_, exists := ns.users[request.User]
	if !exists {
		ns.users[request.User] = domain.NewNotificationRateLimit()
	}

	err := ns.users[request.User].RateLimit(request.Type)
	if err != nil {
		ns.logger.Error(ctx, "error in NotificationService#SendEmail: ", err)
		return err
	}

	ns.logger.Info(ctx, fmt.Sprintf("Email sent. User: %s, Type: %s, Message: %s", request.User, request.Type, request.Message))
	return nil
}

func NewNotification() Notification {
	logger := logs.New("Notification Service")
	return &notification{
		logger: logger,
		users:  make(map[string]*domain.NotificationRateLimit),
	}
}
