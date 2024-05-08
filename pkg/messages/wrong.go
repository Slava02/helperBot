// Package messages contains all messages to be displayed.
package messages

import (
	"cleanBot/models"
)

// wrong returns wrong command or message.
func Wrong(user *models.User) string {
	return "Введите корректное сообщение или команду."
}
