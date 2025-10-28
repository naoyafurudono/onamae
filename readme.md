# onamae

マステでお名前シール作ったら無限です。
テンプレをもとにアイコンと名前入れれば無限お名前シール作れるツールです。SUZURIでマスキングテープを作るために最適な画像を生成します。

以下のような感じに生成できます。

<img width="1181" height="425" alt="donokun " src="https://github.com/user-attachments/assets/ec5a995b-6187-465e-a421-7ae1f3113d8e" />

SUZURIで以下のようにマスキングテープを作るといい感じです。
<iframe height="162" width="375" src="https://suzuri.jp/fnaoya/18840873/masking-tape/15mm/white/embed"></iframe>

## 機能

- SUZURI推奨サイズ（1181×425px）のマスキングテープ用画像を生成
- 名前のみ、または名前+アイコンの2パターンに対応
- シンプルなCLIで簡単に生成可能

## インストール

```bash
go install github.com/naoyafurudono/onamae/cmd/onamae@latest
```

## 使い方

### 基本的な使い方（名前のみ）

```bash
./onamae generate --name "田中太郎"
```

### アイコン付き

```bash
./onamae generate --name "山田花子" --icon icon.png --output hanako.png
```

### オプション

- `--name`: お名前（必須）
- `--icon`: アイコン画像のパス（オプション）
- `--template`: テンプレート画像のパス（デフォルト: `assets/templates/default.png`）
- `--output`: 出力ファイルのパス（デフォルト: `output.png`）

## プロジェクト構成

```
onamae/
├── cmd/
│   ├── onamae/              # CLIツール本体
│   ├── generate-template/   # テンプレート生成スクリプト
│   └── test-generator/      # テスト用スクリプト
├── pkg/
│   └── generator/           # 画像生成ロジック
├── assets/
│   ├── templates/           # テンプレート画像
│   └── fonts/               # フォントファイル
└── readme.md
```

## 今後の予定

- [ ] SUZURI API連携
- [ ] 複数テンプレート対応
- [ ] カラー・フォントのカスタマイズ機能
