package translator

import (
	"fmt"
	"kz_bot/pkg/translator/translategooglefree"
	"os"
)

//	func Translate1(text string) {
//		info := whatlanggo.Detect(text)
//
//		language := info.Lang.String()
//		confidence := info.Confidence
//
//		result := fmt.Sprintf("Определенный язык: %s, Уверенность: %.2f\n", language, confidence)
//
//		saveToFile("output.txt", result+": "+text)
//
// }
func saveToFile(filename string, lines string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.WriteString(lines + "\n")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("Строки сохранены в файл: %s\n", filename)
}
func Translate(text string) {
	result, _ := translategooglefree.Translate(text, "auto", "ru")
	if text != result {
		saveToFile("googleTest.txt", fmt.Sprintf("original: %s \n translate: %s", text, result))
	}
	//fmt.Println(result)
}
func TranslateAnswer(text, langTarget string) string {
	if langTarget == "ua" {
		langTarget = "uk"
	}
	result, _ := translategooglefree.Translate(text, "auto", langTarget)
	return result
}
