package utils

import "regexp"

// Функция для детектирования запрещенных символов

func HasForbiddenChars(input string) bool {
	// Определяем регулярное выражение для запрещенных символов
	// Список символов: < > ' " ; & | / \ $ # ( ) { } * = @ % \r \n
	forbiddenPattern := `[<>';&/\\$#{}@%\r\n]`

	// Компилируем регулярное выражение
	re := regexp.MustCompile(forbiddenPattern)

	// Проверяем, содержится ли хотя бы один запрещенный символ
	return re.MatchString(input)
}
