// Package messages contains all messages to be displayed.
package messages

import (
	"fmt"

	"github.com/Slava02/helperBot/models"
)

// Information returns user's information message.
func Information(user *models.User) string {
	return fmt.Sprintf("Telegram ID: %d\nНикнейм: %s\nИмя: %s\nФамилия: %s\n",
		user.TelegramID,
		user.Username,
		user.FirstName,
		user.LastName,
	)
}
