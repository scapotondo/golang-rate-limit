package domain_test

import (
	"golang-rate-limit/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateLimitSuccess(t *testing.T) {
	notificationRateLimit := domain.NewNotificationRateLimit()

	testCases := map[string]int{
		"news":      1,
		"status":    2,
		"marketing": 3,
	}

	for notificationType, occurencies := range testCases {
		var err error
		t.Run("test rate limit", func(t *testing.T) {
			for i := 0; i < occurencies; i++ {
				err = notificationRateLimit.RateLimit(notificationType)
			}
			assert.NoError(t, err)
		})
	}
}

func TestRateLimitFails(t *testing.T) {
	notificationRateLimit := domain.NewNotificationRateLimit()

	testCases := map[string]int{
		"news":      2,
		"status":    3,
		"marketing": 4,
	}

	for notificationType, occurencies := range testCases {
		var err error
		t.Run("test rate limit", func(t *testing.T) {
			for i := 0; i < occurencies; i++ {
				err = notificationRateLimit.RateLimit(notificationType)
			}
			assert.Error(t, err)
		})
	}
}
