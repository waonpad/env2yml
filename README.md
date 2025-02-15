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

#### Darwin ARM64 用にビルドする場合

```bash
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o build/env2yml-darwin-arm64 main.go; strip build/env2yml-darwin-arm64;
```

#### Linux AMD64 用にビルドする場合

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o build/env2yml-linux-amd64 main.go; strip build/env2yml-linux-amd64;
```

#### ビルド対象のOSとアーキテクチャの一覧を確認できる

```bash
go tool dist list
```
