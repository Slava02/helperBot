package messages

import (
	"fmt"
	"github.com/Slava02/helperBot/models"
)

func Media(user *models.User) string {
	return "Что вас интересует:"
}

func MediaOptions() string {
	return fmt.Sprintf("Удалить или сохранить?")
}

func NoData() string {
	return fmt.Sprintf("В таблице пока нет данных")
}

func AlreadySaved() string {
	return fmt.Sprintf("Уже сохранено")
}
