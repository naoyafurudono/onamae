package main

import (
	"image/color"
	"log"

	"github.com/fogleman/gg"
)

func main() {
	// マスキングテープのサイズ（SUZURI推奨サイズ: 幅15mm × 全長5m）
	const width = 1181
	const height = 425

	dc := gg.NewContext(width, height)

	// 背景色（オフホワイト）
	dc.SetColor(color.RGBA{R: 250, G: 250, B: 248, A: 255})
	dc.Clear()

	// 細いボーダーで枠取り（上下左右）
	borderColor := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	borderWidth := 2.0

	dc.SetColor(borderColor)
	dc.SetLineWidth(borderWidth)

	// 上ボーダー
	dc.DrawLine(0, borderWidth/2, float64(width), borderWidth/2)
	dc.Stroke()

	// 下ボーダー
	dc.DrawLine(0, float64(height)-borderWidth/2, float64(width), float64(height)-borderWidth/2)
	dc.Stroke()

	// 左ボーダー
	dc.DrawLine(borderWidth/2, 0, borderWidth/2, float64(height))
	dc.Stroke()

	// 右ボーダー
	dc.DrawLine(float64(width)-borderWidth/2, 0, float64(width)-borderWidth/2, float64(height))
	dc.Stroke()

	// テンプレート画像を保存
	if err := dc.SavePNG("assets/templates/default.png"); err != nil {
		log.Fatal(err)
	}

	log.Println("テンプレート画像を生成しました: assets/templates/default.png")
}
