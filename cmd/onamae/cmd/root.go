package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "onamae",
	Short: "お名前シール生成ツール",
	Long: `onamae - お名前シール生成ツール

マスキングテープ用のお名前シールを簡単に生成できます。
SUZURI推奨サイズ（1181×425px）に対応しています。`,
	Example: `  # 名前のみ
  onamae generate --name 太郎

  # アイコン付き
  onamae generate --name 花子 --icon icon.png

  # 出力先を指定
  onamae generate --name 太郎 --output tanaka.png`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
