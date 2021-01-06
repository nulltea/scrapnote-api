package service

import "context"

type MailService interface {
	SendEmailConfirmation(ctx context.Context, email, callbackURL string) error
	SendResetPassword(ctx context.Context, email, callbackURL string) error
	SendNotification(ctx context.Context, email, notificationContent string) error
}