// Package messages contains all messages to be displayed.
package messages

import (
	"cleanBot/models"
)

// Help returns a help message.
func Help(user *models.User) string {
	return "Это Бот с чистой архитектурой на Golang!"
}
