package utils

import "sort"

// Функция для вычисления медианы массива чисел
func Median(numbers []int) int {
	// Сортировать массив
	sort.Ints(numbers)
	// Вычислить медиану в зависимости от длины массива
	length := len(numbers)
	if length%2 == 0 {
		// Если длина чётная, усреднить два средних элемента
		mid := length / 2
		median := (numbers[mid-1] + numbers[mid]) / 2.0
		return median
	} else {
		// Если длина нечётная, взять средний элемент
		mid := length / 2
		median := numbers[mid]
		return median
	}
}
