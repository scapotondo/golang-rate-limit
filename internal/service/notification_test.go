package service_test

import (
	"context"
	"errors"
	"golang-rate-limit/internal/resources"
	"golang-rate-limit/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

type sendEmailCase struct {
	testName      string
	occurencies   int
	request       resources.NotificationRequest
	expectedError error
}

func TestSendEmail(t *testing.T) {
	testCases := []sendEmailCase{
		{
			testName:    "Successfull status case",
			occurencies: 2,
			request: resources.NotificationRequest{
				Type:    "status",
				User:    "user-1",
				Message: "message",
			},
		},
		{
			testName:    "Error status case",
			occurencies: 3,
			request: resources.NotificationRequest{
				Type:    "status",
				User:    "user-1",
				Message: "message",
			},
			expectedError: errors.New("error"),
		},
		{
			testName:    "Successfull news case",
			occurencies: 1,
			request: resources.NotificationRequest{
				Type:    "news",
				User:    "user-1",
				Message: "message",
			},
		},
		{
			testName:    "Error status case",
			occurencies: 2,
			request: resources.NotificationRequest{
				Type:    "news",
				User:    "user-1",
				Message: "message",
			},
			expectedError: errors.New("error"),
		},
		{
			testName:    "Successfull marketing case",
			occurencies: 1,
			request: resources.NotificationRequest{
				Type:    "marketing",
				User:    "user-1",
				Message: "message",
			},
		},
		{
			testName:    "Error status case",
			occurencies: 4,
			request: resources.NotificationRequest{
				Type:    "marketing",
				User:    "user-1",
				Message: "message",
			},
			expectedError: errors.New("error"),
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {

			notificationService := service.NewNotification()
			for i := 0; i < test.occurencies; i++ {
				err := notificationService.SendEmail(context.Background(), test.request)
				if err != nil {
					require.Equal(t, test.expectedError, err)
				}
			}
		})
	}
}
