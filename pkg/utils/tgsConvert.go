package utils

import (
	"bytes"
	"fmt"
	"github.com/Benau/tgsconverter/libtgsconverter"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

func Convert(url string) (string, []byte) {
	// Скачиваем файл по URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fileName := filepath.Base(url)
	// Читаем содержимое файла
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	data := buf.Bytes()
	name := filepath.Base(fileName)
	if filepath.Ext(name) == ".tgs" {
		name += ".webp"
	}
	if strings.HasSuffix(name, ".tgs.webp") {
		maybeConvertTgs(&name, &data)
	} else if strings.HasSuffix(name, ".webp") {
		//d.maybeConvertWebp(&name, &data)
	}

	return name, data
}
func maybeConvertWebp(name *string, data *[]byte) {
	//err := helper.ConvertWebPToPNG(data)
	//if err != nil {
	//	b.Log.Errorf("conversion failed: %v", err)
	//} else {
	//	*name = strings.Replace(*name, ".webp", ".png", 1)
	//}
}
func maybeConvertTgs(name *string, data *[]byte) {

	err := ConvertTgsToX(data, "png")
	if err != nil {
		fmt.Errorf("conversion failed: %v", err)
	} else {
		*name = strings.Replace(*name, "tgs.webp", "png", 1)
	}
}
func ConvertTgsToX(data *[]byte, outputFormat string) error {
	options := libtgsconverter.NewConverterOptions()
	options.SetExtension(outputFormat)
	blob, err := libtgsconverter.ImportFromData(*data, options)
	if err != nil {
		return fmt.Errorf("failed to run libtgsconverter.ImportFromData: %s", err.Error())
	}

	*data = blob
	return nil
}
