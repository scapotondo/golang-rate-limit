package gin_handlers_test

import (
	gin_handlers "golang-rate-limit/internal/gin-handlers"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestRateLimitSuccess(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	notificationRateLimit := gin_handlers.NewNotificationRateLimit()

	testCases := map[string]int{
		"news":      1,
		"status":    2,
		"marketing": 3,
	}

	for notificationType, occurencies := range testCases {
		t.Run("test rate limit", func(t *testing.T) {

			c.Params = []gin.Param{
				{
					Key:   "type",
					Value: notificationType,
				},
			}

			for i := 0; i < occurencies; i++ {
				notificationRateLimit.RateLimit(c)
			}

			assert.Equal(t, false, c.IsAborted())
		})
	}
}

func TestRateLimitFails(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	notificationRateLimit := gin_handlers.NewNotificationRateLimit()

	testCases := map[string]int{
		"news":      2,
		"status":    3,
		"marketing": 4,
	}

	for notificationType, occurencies := range testCases {
		t.Run("test rate limit", func(t *testing.T) {

			c.Params = []gin.Param{
				{
					Key:   "type",
					Value: notificationType,
				},
			}

			for i := 0; i < occurencies; i++ {
				notificationRateLimit.RateLimit(c)
			}

			assert.Equal(t, true, c.IsAborted())
		})
	}
}
