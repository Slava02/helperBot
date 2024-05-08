// Package messages contains all messages to be displayed.
package messages

import (
	"fmt"

	"cleanBot/models"
)

// Main returns main menu's message.
func Main(user *models.User) string {
	return fmt.Sprintf("Добро пожаловать %s.", user.FirstName)
}
