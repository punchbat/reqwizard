package utils

func RemoveElementByIndex[T any](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}