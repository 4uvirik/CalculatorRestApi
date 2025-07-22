package calc

import "testing"

func TestMultiplyNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Пустой слайс", nil, 0},
		{"Нулевое значение", []int{0}, 0},
		{"положительные числа", []int{2, 3, 5}, 30},
		{"отрицательные числа", []int{-2, -3, -5}, -30},
		{"разные числа", []int{-2, 0, 5}, -10},
		{"одно число", []int{351}, 351},
		{"большой слайс", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
		{"большие числа", []int{111111, 555555, 999999}, 1666665},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SumNumbers(tt.input)
			if actual != tt.expected {
				t.Errorf("SumNumbers %v = %d ; want %d", tt.input, actual, tt.expected)
			}
		})
	}
}
