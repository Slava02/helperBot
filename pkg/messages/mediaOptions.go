package messages

import (
	"fmt"
)

func MediaOptions() string {
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

func NoData() string {
	return fmt.Sprintf("В таблице пока нет данных")
}

func PickRandomErr() string {
	return fmt.Sprintf("Ошибка при выборе случайного объекта, попробуйте еще раз")
}
