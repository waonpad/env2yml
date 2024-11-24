# .env ファイルを yml 形式に変換する

## ビルド済みバイナリを利用する

build ディレクトリに一部の環境向けのビルド済みバイナリが格納されている

### 実行

```bash
./build/env2yml-darwin-arm64 ./fixtures/.env.test
```

## 自分でビルドする

### Go のインストール

```
brew install go
```

### ビルド

Darwin ARM64 用にビルドする場合

```bash
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o build/env2yml-darwin-arm64 main.go; strip build/env2yml-darwin-arm64;
```

#### upx で更に圧縮する場合

```bash
brew install upx
```

```bash
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o build/env2yml-darwin-arm64 main.go; strip build/env2yml-darwin-arm64; upx --lzma build/env2yml-darwin-arm64
```

以下コマンドで OS とアーキテクチャの一覧を確認できる

```bash
go tool dist list
```
