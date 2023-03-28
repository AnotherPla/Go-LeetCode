package main

func MinArray(numbers []int) int {
	var left, right = 0, len(numbers) - 1
	var mid = ((right - left) >> 1) + left
	for left < right {
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
		mid = ((right - left) >> 1) + left
	}
	return numbers[mid]
}
