# .env ファイルを yml 形式に変換する

## ビルド済みバイナリを利用する

build ディレクトリに一部の環境向けのビルド済みバイナリが格納されている

### 実行

```bash
./build/env-2-yml-darwin-arm64 ./fixtures/.env.test
```

## 自分でビルドする

### Go のインストール

```
brew install go
```

### ビルド

Linux ARM64 用にビルドする場合

```bash
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build
```

以下コマンドで OS とアーキテクチャの一覧を確認できる

```bash
go tool dist list
```
