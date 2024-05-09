// Package messages contains all messages to be displayed.
package messages

import (
	"github.com/Slava02/helperBot/models"
)

// wrong returns wrong command or message.
func Wrong(user *models.User) string {
	return "Введите корректное сообщение или команду."
}
