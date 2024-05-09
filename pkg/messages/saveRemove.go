package messages

import (
	"fmt"
	"github.com/Slava02/helperBot/models"
)

func SaveOrRemove(user *models.User) string {
	return fmt.Sprintf("Удалить или сохранить?")
}
