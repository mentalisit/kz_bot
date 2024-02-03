package imageGenerator

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"net/http"
	"strconv"
)

var user map[int][]int

func GetLevel(i int) string {
	l := user[i]
	if len(l) > 0 {
		level := strconv.Itoa(l[0])
		return level
	}
	return "0"
}
func addModulesLevel(dc *gg.Context) {
	y := float64(305)
	x := float64(100)
	x2 := x + 90
	x3 := x2 + 90
	x4 := x3 + 90
	x5 := x4 + 185
	x6 := x5 + 90
	x7 := x6 + 90
	x8 := x7 + 90
	//transport1
	dc.DrawStringAnchored(GetLevel(401), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(402), x2, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(413), x3, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(404), x4, y, 0, 0.5)
	//mainer1
	dc.DrawStringAnchored(GetLevel(501), x5, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(511), x6, y, 0, 0.5) ////
	dc.DrawStringAnchored(GetLevel(512), x7, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(504), x8, y, 0, 0.5)
	///
	y += 70
	//transport2
	dc.DrawStringAnchored(GetLevel(608), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(405), x2, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(406), x3, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(603), x4, y, 0, 0.5)
	//mainer2
	dc.DrawStringAnchored(GetLevel(508), x5, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(503), x6, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(507), x7, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(505), x8, y, 0, 0.5)
	///
	y += 70
	//transport3
	dc.DrawStringAnchored(GetLevel(412), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(411), x2, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(414), x3, y, 0, 0.5)
	//mainer3
	dc.DrawStringAnchored(GetLevel(510), x5, y, 0, 0.5) ////
	dc.DrawStringAnchored(GetLevel(513), x6, y, 0, 0.5) ////

	////
	y += 110
	//weapon1
	dc.DrawStringAnchored(GetLevel(203), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(204), x2, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(202), x3, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(205), x4, y, 0, 0.5)
	//Shield1
	dc.DrawStringAnchored(GetLevel(301), x5, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(302), x6, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(303), x7, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(304), x8, y, 0, 0.5)
	///
	y += 70
	//weapon2
	dc.DrawStringAnchored(GetLevel(206), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(207), x2, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(208), x3, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(209), x4, y, 0, 0.5)
	//Shield2
	dc.DrawStringAnchored(GetLevel(306), x5, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(305), x6, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(307), x7, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(308), x8, y, 0, 0.5)
	y += 70
	//weapon3
	dc.DrawStringAnchored(GetLevel(210), x, y, 0, 0.5)
	////
	y += 110
	//support1
	dc.DrawStringAnchored(GetLevel(601), x, y, 0, 0.5)  //90
	dc.DrawStringAnchored(GetLevel(625), x2, y, 0, 0.5) /////
	dc.DrawStringAnchored(GetLevel(609), x3, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(602), x4, y, 0, 0.5)
	//drone1
	dc.DrawStringAnchored(GetLevel(901), x5, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(902), x6, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(904), x7, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(624), x8, y, 0, 0.5)
	///
	y += 70
	//support2
	dc.DrawStringAnchored(GetLevel(626), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(614), x2, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(615), x3, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(616), x4, y, 0, 0.5)
	//drone2
	dc.DrawStringAnchored(GetLevel(905), x5, y, 0, 0.5) ////
	dc.DrawStringAnchored(GetLevel(906), x6, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(907), x7, y, 0, 0.5)
	///
	y += 70
	//support3
	dc.DrawStringAnchored(GetLevel(617), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(618), x2, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(619), x3, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(622), x4, y, 0, 0.5) ///////?
	y += 70
	//support4
	dc.DrawStringAnchored(GetLevel(621), x, y, 0, 0.5) //90
	dc.DrawStringAnchored(GetLevel(623), x2, y, 0, 0.5)
	//level
	dc.DrawStringAnchored(GetLevel(101), x5, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(103), x6, y, 0, 0.5)
	dc.DrawStringAnchored(GetLevel(102), x7, y, 0, 0.5)
}
func addAvatars(dc *gg.Context, AvatarURL string, centerX, centerY int) {
	// Загружаем изображение аватара
	response, err := http.Get(AvatarURL)
	if err != nil {
		//fmt.Println(err)
		return
	}

	avatar, _, err := image.Decode(response.Body)

	// Изменяем размер изображения аватара
	avatar = resize.Resize(185, 185, avatar, resize.Lanczos3)

	// Определяем радиус круга
	radius := 270 / 3 //270
	// Рисуем круг с изображением
	dc.ResetClip()
	dc.DrawCircle(float64(centerX), float64(centerY), float64(radius))
	dc.Clip()
	dc.DrawImageAnchored(avatar, centerX, centerY, 0.5, 0.5)
	response.Body.Close()
	return
}
func imageToBytesReader(img image.Image) (*bytes.Reader, error) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}
