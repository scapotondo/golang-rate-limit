package gin_handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type NotificationRateLimit struct {
	statusRateLimiter    *rate.Limiter
	newsRateLimiter      *rate.Limiter
	marketingRateLimiter *rate.Limiter
}

func (nrl *NotificationRateLimit) RateLimit(ctx *gin.Context) {
	param, _ := ctx.Params.Get("type")

	// This variable is false by default to only allow news, status and marketing mails
	allow := false
	switch param {
	case "news":
		allow = nrl.newsRateLimiter.Allow()
	case "status":
		allow = nrl.statusRateLimiter.Allow()
	case "marketing":
		allow = nrl.marketingRateLimiter.Allow()
	}

	if !allow {
		ctx.Status(http.StatusTooManyRequests)
		ctx.Abort()
	}

	ctx.Next()
}

func NewNotificationRateLimit() *NotificationRateLimit {
	return &NotificationRateLimit{
		statusRateLimiter:    rate.NewLimiter(rate.Every(time.Minute), 2),
		newsRateLimiter:      rate.NewLimiter(rate.Every(24*time.Hour), 1),
		marketingRateLimiter: rate.NewLimiter(rate.Every(time.Hour), 3),
	}
}
