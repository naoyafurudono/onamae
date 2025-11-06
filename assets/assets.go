package assets

import (
	_ "embed"
)

// DefaultFont は埋め込まれたデフォルトフォント（IPAゴシック）
//
//go:embed fonts/ipaexg.ttf
var DefaultFont []byte

// DefaultTemplate は埋め込まれたデフォルトテンプレート
//
//go:embed templates/default.png
var DefaultTemplate []byte
