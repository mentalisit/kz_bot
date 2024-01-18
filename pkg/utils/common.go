package utils

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}
		return nil
	}
	return
}
func RemoveDuplicateElementString(a []string) []string {
	result := make([]string, 0, len(a))
	temp := map[string]struct{}{}
	for _, item := range a {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func RemoveDuplicateElementInt(a []int) []int {
	result := make([]int, 0, len(a))
	temp := map[int]struct{}{}
	for _, item := range a {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func GetRandomColor() string {
	// Генерируем случайные значения для красного, зеленого и синего цветов
	red := rand.Intn(256)
	green := rand.Intn(256)
	blue := rand.Intn(256)

	// Форматируем цвет в HEX
	colorHex := fmt.Sprintf("%02X%02X%02X", red, green, blue)

	return colorHex
}
func ExtractUppercase(input string) string {
	var result string

	for _, char := range input {
		if unicode.IsUpper(char) {
			result += string(char)
		}
	}
	if result == "" {
		return input
	}

	return result
}
