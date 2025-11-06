package generator

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

// Generator はお名前シール生成器
type Generator struct {
	templatePath string
}

// New は新しいGeneratorを作成します
func New(templatePath string) *Generator {
	return &Generator{
		templatePath: templatePath,
	}
}

// GenerateWithName はテンプレートに名前を描画した画像を生成します
func (g *Generator) GenerateWithName(name string, outputPath string) error {
	// テンプレート画像を読み込み
	templateImg, err := gg.LoadImage(g.templatePath)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	// 画像の幅と高さを取得
	bounds := templateImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// 新しいコンテキストを作成してテンプレートを描画
	dc := gg.NewContext(width, height)
	dc.DrawImage(templateImg, 0, 0)

	// フォントを設定（IPAゴシックフォント）
	if err := dc.LoadFontFace("assets/fonts/ipaexg.ttf", 120); err != nil {
		return fmt.Errorf("failed to load font: %w", err)
	}

	// テキストの色を設定（ダークグレー）
	dc.SetColor(color.RGBA{R: 60, G: 60, B: 60, A: 255})

	// テキストを中央に配置
	dc.DrawStringAnchored(name, float64(width)/2, float64(height)/2, 0.5, 0.5)

	// 画像を保存
	if err := dc.SavePNG(outputPath); err != nil {
		return fmt.Errorf("failed to save output: %w", err)
	}

	return nil
}

// GenerateWithNameAndIcon はテンプレートに名前とアイコンを描画した画像を生成します
func (g *Generator) GenerateWithNameAndIcon(name string, iconPath string, outputPath string) error {
	// テンプレート画像を読み込み
	templateImg, err := gg.LoadImage(g.templatePath)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	// アイコン画像を読み込み
	iconImg, err := gg.LoadImage(iconPath)
	if err != nil {
		return fmt.Errorf("failed to load icon: %w", err)
	}

	// 画像の幅と高さを取得
	bounds := templateImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// 新しいコンテキストを作成してテンプレートを描画
	dc := gg.NewContext(width, height)
	dc.DrawImage(templateImg, 0, 0)

	// アイコンをリサイズして描画（高さの80%のサイズに）
	iconSize := int(float64(height) * 0.8)
	resizedIcon := resizeImage(iconImg, iconSize, iconSize)

	// アイコンを左側に配置
	iconX := height / 10
	iconY := (height - iconSize) / 2
	dc.DrawImage(resizedIcon, iconX, iconY)

	// フォントを設定（IPAゴシックフォント）
	if err := dc.LoadFontFace("assets/fonts/ipaexg.ttf", 120); err != nil {
		return fmt.Errorf("failed to load font: %w", err)
	}

	// テキストの色を設定（ダークグレー）
	dc.SetColor(color.RGBA{R: 60, G: 60, B: 60, A: 255})

	// テキストをアイコンの右側に配置
	textX := float64(iconX + iconSize + 50)
	textY := float64(height) / 2
	dc.DrawStringAnchored(name, textX, textY, 0, 0.5)

	// 画像を保存
	if err := dc.SavePNG(outputPath); err != nil {
		return fmt.Errorf("failed to save output: %w", err)
	}

	return nil
}

// GenerateWithTwoPatterns はテンプレートを2分割して、それぞれに名前とアイコンを描画した画像を生成します
func (g *Generator) GenerateWithTwoPatterns(name1, icon1, name2, icon2, outputPath string) error {
	// テンプレート画像を読み込み
	templateImg, err := gg.LoadImage(g.templatePath)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	// 画像の幅と高さを取得
	bounds := templateImg.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	halfWidth := width / 2

	// 新しいコンテキストを作成（テンプレートと同じサイズ）
	dc := gg.NewContext(width, height)
	dc.DrawImage(templateImg, 0, 0)

	// 左半分に1つ目のパターンを描画
	if err := g.drawVerticalPattern(dc, name1, icon1, 0, halfWidth, height); err != nil {
		return fmt.Errorf("failed to draw first pattern: %w", err)
	}

	// 右半分に2つ目のパターンを描画
	if err := g.drawVerticalPattern(dc, name2, icon2, halfWidth, width-halfWidth, height); err != nil {
		return fmt.Errorf("failed to draw second pattern: %w", err)
	}

	// 画像を保存
	if err := dc.SavePNG(outputPath); err != nil {
		return fmt.Errorf("failed to save output: %w", err)
	}

	return nil
}

// drawVerticalPattern は指定された領域に名前とアイコンを縦配置で描画します
func (g *Generator) drawVerticalPattern(dc *gg.Context, name, iconPath string, offsetX, areaWidth, height int) error {
	// アイコンが指定されている場合
	if iconPath != "" {
		// アイコン画像を読み込み
		iconImg, err := gg.LoadImage(iconPath)
		if err != nil {
			return fmt.Errorf("failed to load icon: %w", err)
		}

		// アイコンのサイズを決定（高さの75%）
		iconSize := int(float64(height) * 0.75)
		resizedIcon := resizeImage(iconImg, iconSize, iconSize)

		// 縦配置（上にアイコン、下に名前）
		// アイコンを上部中央に配置
		iconX := offsetX + (areaWidth-iconSize)/2
		iconY := int(float64(height) * 0.05) // 上から5%の位置
		dc.DrawImage(resizedIcon, iconX, iconY)

		// フォントを設定
		if err := dc.LoadFontFace("assets/fonts/ipaexg.ttf", 70); err != nil {
			return fmt.Errorf("failed to load font: %w", err)
		}

		// テキストの色を設定
		dc.SetColor(color.RGBA{R: 60, G: 60, B: 60, A: 255})

		// テキストをアイコンの下、中央に配置
		textX := float64(offsetX) + float64(areaWidth)/2
		textY := float64(iconY + iconSize + 55) // アイコンの下、55px空けて配置
		dc.DrawStringAnchored(name, textX, textY, 0.5, 0)
	} else {
		// アイコンなしの場合、名前だけを中央に配置
		// フォントを設定
		if err := dc.LoadFontFace("assets/fonts/ipaexg.ttf", 70); err != nil {
			return fmt.Errorf("failed to load font: %w", err)
		}

		// テキストの色を設定
		dc.SetColor(color.RGBA{R: 60, G: 60, B: 60, A: 255})

		// テキストを領域の中央に配置
		textX := float64(offsetX) + float64(areaWidth)/2
		textY := float64(height) / 2
		dc.DrawStringAnchored(name, textX, textY, 0.5, 0.5)
	}

	return nil
}

// resizeImage は画像をリサイズします
func resizeImage(img image.Image, width, height int) image.Image {
	dc := gg.NewContext(width, height)
	// スケールを先に設定してから画像を描画
	scaleX := float64(width) / float64(img.Bounds().Dx())
	scaleY := float64(height) / float64(img.Bounds().Dy())
	dc.Scale(scaleX, scaleY)
	dc.DrawImage(img, 0, 0)
	return dc.Image()
}
