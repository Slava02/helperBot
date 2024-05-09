package messages

import (
	"fmt"
)

func SaveOrRemove() string {
	return fmt.Sprintf("Удалить или сохранить?")
}

func SaveWhat() string {
	return fmt.Sprintf("Введите названия: ")
}

func SaveSuccess() string {
	return fmt.Sprintf("Успешно сохранено!")
}

func SaveError() string {
	return fmt.Sprintf("Не удалось сохранить, попробуйте еще раз:")
}

func Remove() string {
	return fmt.Sprintf("Введите номер: ")
}
