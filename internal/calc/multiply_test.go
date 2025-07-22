package calc

import "testing"

func TestMultiplyNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Пустой слайс", nil, 1},
		{"Нулевое значение", []int{0}, 0},
		{"положительные числа", []int{2, 3, 5}, 30},
		{"отрицательные числа", []int{-2, -3, -5}, -30},
		{"разные числа", []int{-2, 1, 5}, -10},
		{"умножение на 0", []int{-2, 0, 5}, 0},
		{"одно число", []int{351}, 351},
		{"большой слайс", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 479001600},
		{"большие числа", []int{111111, 555555, 999999}, 61728209876728395},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MultiplyNumbers(tt.input)
			if actual != tt.expected {
				t.Errorf("MultiplyNumbers %v = %d ; want %d", tt.input, actual, tt.expected)
			}
		})
	}
}
