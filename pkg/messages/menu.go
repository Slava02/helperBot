// Package messages contains all messages to be displayed.
package messages

import (
	"fmt"

	"github.com/Slava02/helperBot/models"
)

// Main returns main menu's message.
func Main(user *models.User) string {
	return fmt.Sprintf("Добро пожаловать %s.", user.FirstName)
}
