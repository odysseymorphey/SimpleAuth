package services

import (
	"github.com/sirupsen/logrus"
)

func sendEmailNotification(email string) {
	logrus.Info("Email notification sent to: ", email)
}
