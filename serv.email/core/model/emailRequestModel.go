package model

type EmailRequest struct {
	Email       string
	CallbackURL string
	Content     string
}