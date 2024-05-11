package messages

import (
	"fmt"
)

func SaveWhat() string {
	return fmt.Sprintf("Введите названия: ")
}

func SaveSuccess() string {
	return fmt.Sprintf("Успешно сохранено!")
}

func SaveError() string {
	return fmt.Sprintf("Не удалось сохранить, попробуйте еще раз:")
}

func RemoveWhat() string {
	return fmt.Sprintf("Выберите объект, который вы хотите удалить:\n")
}

func RemoveSuccess() string {
	return fmt.Sprintf("Данные успешо удалены!")
}
