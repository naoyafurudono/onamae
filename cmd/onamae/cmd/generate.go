package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/naoyafurudono/onamae/pkg/generator"
	"github.com/spf13/cobra"
)

var (
	name     string
	icon     string
	template string
	output   string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "お名前シールを生成",
	Long: `テンプレートに名前（とアイコン）を合成してお名前シールを生成します。

生成される画像はSUZURI推奨サイズ（1181×425px）です。`,
	Example: `  # 名前のみ
  onamae generate --name 太郎

  # アイコン付き
  onamae generate --name 花子 --icon ./icon.png

  # 出力先を指定
  onamae generate --name 太郎 --output ./output/tanaka.png`,
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&name, "name", "n", "", "お名前（必須）")
	generateCmd.Flags().StringVarP(&icon, "icon", "i", "", "アイコン画像のパス（オプション）")
	generateCmd.Flags().StringVarP(&template, "template", "t", "assets/templates/default.png", "テンプレート画像のパス")
	generateCmd.Flags().StringVarP(&output, "output", "o", "output.png", "出力ファイルのパス")

	generateCmd.MarkFlagRequired("name")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// バリデーション
	if name == "" {
		return fmt.Errorf("名前を指定してください（--name）")
	}

	// テンプレートファイルの存在確認
	if _, err := os.Stat(template); os.IsNotExist(err) {
		return fmt.Errorf("テンプレートファイルが見つかりません: %s", template)
	}

	// アイコンファイルの存在確認（指定されている場合）
	if icon != "" {
		if _, err := os.Stat(icon); os.IsNotExist(err) {
			return fmt.Errorf("アイコンファイルが見つかりません: %s", icon)
		}
	}

	// Generator作成
	gen := generator.New(template)

	// 生成処理
	color.Cyan("🎨 お名前シールを生成中...")
	var err error
	if icon != "" {
		err = gen.GenerateWithNameAndIcon(name, icon, output)
	} else {
		err = gen.GenerateWithName(name, output)
	}

	if err != nil {
		color.Red("✗ 生成に失敗しました: %v", err)
		return err
	}

	// 成功メッセージ
	color.Green("✓ お名前シールを生成しました: %s", output)

	// 詳細情報
	fmt.Println()
	color.Yellow("📋 生成情報:")
	fmt.Printf("   名前: %s\n", name)
	if icon != "" {
		fmt.Printf("   アイコン: %s\n", icon)
	}
	fmt.Printf("   出力: %s\n", output)
	fmt.Printf("   サイズ: 1181×425px (SUZURI推奨)\n")

	return nil
}
