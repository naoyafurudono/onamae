package main

import (
	"image/color"
	"log"
	"math"

	"github.com/fogleman/gg"
)

func main() {
	// シンプルな円形アイコンを作成
	const size = 300
	dc := gg.NewContext(size, size)

	// 背景を透明に
	dc.SetColor(color.RGBA{R: 0, G: 0, B: 0, A: 0})
	dc.Clear()

	// 円を描画（水色）
	dc.DrawCircle(float64(size)/2, float64(size)/2, float64(size)/2-10)
	dc.SetColor(color.RGBA{R: 100, G: 200, B: 255, A: 255})
	dc.Fill()

	// 星マーク（白）
	dc.SetColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	centerX, centerY := float64(size)/2, float64(size)/2
	radius := float64(size) / 4

	// 5点星を描画
	for i := 0; i < 5; i++ {
		angle := float64(i) * 4 * math.Pi / 5
		x := centerX + radius*math.Cos(angle-math.Pi/2)
		y := centerY + radius*math.Sin(angle-math.Pi/2)
		if i == 0 {
			dc.MoveTo(x, y)
		} else {
			dc.LineTo(x, y)
		}
	}
	dc.ClosePath()
	dc.Fill()

	if err := dc.SavePNG("test_icon.png"); err != nil {
		log.Fatal(err)
	}

	log.Println("テストアイコンを生成しました: test_icon.png")
}
