package imageGenerator

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"log"
)

func GenerateUser(avatarURL, corpAvararUrl, nikName, corporation string, tech map[int][]int) *bytes.Reader {
	//avatarURL := "https://cdn.discordapp.com/avatars/582882137842122773/c41d183fba78b9f49bb590a1fc8e33a9.png"
	//corpAvararUrl := "https://cdn.discordapp.com/icons/716771579278917702/d539d5d8f722b1402e323c57a7a48317.png"
	//nikName := "My NickName"
	//corporation := "My Corporation"
	user = tech
	// Открываем изображение
	im, err := gg.LoadPNG("./config/original2.png")
	if err != nil {
		fmt.Println(err)
	}

	// Создаем новый контекст для рисования
	dc := gg.NewContextForImage(im)

	// Устанавливаем параметры шрифта
	err = dc.LoadFontFace("./config/font.ttf", 32)
	if err != nil {
		fmt.Println(err)
	}

	// Устанавливаем цвет текста
	dc.SetColor(color.White)
	dc.DrawStringAnchored(nikName, 450, 100, 0.5, 0.5)
	dc.DrawStringAnchored(corporation, 450, 150, 0.5, 0.5)

	// Рисуем текст на изображении 1
	addModulesLevel(dc)
	addAvatars(dc, avatarURL, 125, 125)
	addAvatars(dc, corpAvararUrl, 760, 125)

	reader, err := imageToBytesReader(dc.Image())
	if err != nil {
		log.Fatal(err)
	}
	return reader
	//// Сохраняем изображение
	//err = dc.SavePNG("output.png")
	//if err != nil {
	//	fmt.Println(err)
	//}
}
