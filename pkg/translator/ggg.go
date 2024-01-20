package translator

//import (
//	"fmt"
//	"github.com/bregydoc/gtranslate"
//	"golang.org/x/text/language"
//)
//
//func Gmain() {
//	textToTranslate := "привіт як ти ?"
//
//	translatedText, err := gtranslate.Translate(textToTranslate, language.English, language.Russian)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	params, err := gtranslate.TranslateWithParams(textToTranslate, gtranslate.TranslationParams{
//		From: "auto",
//		To:   "ru",
//	})
//	if err != nil {
//		return
//	}
//	fmt.Printf("tt %+v\n", params)
//	fmt.Printf("Original text: %s\nTranslated text: %s\n", textToTranslate, translatedText)
//}
