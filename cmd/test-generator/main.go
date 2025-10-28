package main

import (
	"log"

	"github.com/naoyafurudono/onamae/pkg/generator"
)

func main() {
	gen := generator.New("assets/templates/default.png")

	// 名前のみのテスト
	if err := gen.GenerateWithName("太郎", "output_name_only.png"); err != nil {
		log.Fatal(err)
	}
	log.Println("名前のみの画像を生成しました: output_name_only.png")

	// 名前とアイコンのテスト
	if err := gen.GenerateWithNameAndIcon("花子", "test_icon.png", "output_with_icon.png"); err != nil {
		log.Fatal(err)
	}
	log.Println("名前とアイコンの画像を生成しました: output_with_icon.png")
}
