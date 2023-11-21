package resources

type NotificationRequest struct {
	Type    string `form:"type" example:"news" enums:"news,status,marketing"`
	User    string `form:"user" example:"user-1" binding:"required"`
	Message string `form:"message" example:"some message" binding:"required"`
}
