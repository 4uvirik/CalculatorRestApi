package calc

import (
	"errors"
	"strconv"
	"strings"
)

func SumForPostString(jsonPost string) (int, error) {
	digits, err := convertStringToInt(jsonPost)
	if err != nil {
		return 0, err
	}
	return sumNumbers(digits), nil
}

// Конвертация строку в числа
func convertStringToInt(input string) ([]int, error) {
	input = strings.ReplaceAll(input, ",", " ")
	numbers := strings.Fields(input) // Функция разбивает данные по пробелам
	// Проверяем не пустая ли строка
	if len(numbers) == 0 {
		return nil, errors.New("Введите хотя бы два числа\n")
	}

	var number []int
	for _, n := range numbers {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, errors.New("Недопустимое значение, введите целые числа через пробел или запятую\n")
		}
		number = append(number, num)
	}
	return number, nil
}

// Подсчет суммы чисел в слайсе
func sumNumbers(number []int) int {
	sum := 0
	for _, n := range number {
		sum += n
	}
	return sum
}
