package utils

func ConvertArrayToString(arr []int) string {
	var str string
	for i := range arr[:] {
		str += string(rune(arr[i]))
	}
	return str
}

func SquareSum(numbers []int) int {
	var count int
	for i := range numbers[:] {
		count += numbers[i] * numbers[i]
	}
	return count
}
