package service

type MailService interface {
	SendEmailConfirmation(email, callbackURL string) error
	SendResetPassword(email, callbackURL string) error
	SendNotification(email, notificationContent string) error
}