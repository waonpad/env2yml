package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    // 引数が存在するか確認
    if len(os.Args) <= 1 {
        fmt.Println("Please provide the path to the .env file")
        os.Exit(1)
    }

    // 入力された .env ファイルのパスを取得
    envFile := os.Args[1]

    // .env ファイルが存在するか確認
    if _, err := os.Stat(envFile); os.IsNotExist(err) {
        fmt.Println("The specified .env file does not exist")
        os.Exit(1)
    }

    // .env ファイルのファイル名の後ろに .yml を付けたファイル名を生成
    yamlFile := envFile + ".yml"

    // .env ファイルの内容を読み込む
    file, err := os.Open(envFile)
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(1)
    }
    defer file.Close()

    var envYamlContent strings.Builder
    scanner := bufio.NewScanner(file)

	// .env ファイルの内容を読み込み、YAML 形式に変換
    for scanner.Scan() {
        line := scanner.Text()
        // 空文字列またはコメント行はそのまま出力
        if strings.TrimSpace(line) == "" || strings.HasPrefix(strings.TrimSpace(line), "#") {
            envYamlContent.WriteString(line + "\n")
            continue
        }

        // key=value の形式を分割
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            fmt.Printf("Invalid line in .env file: %s\n", line)
            os.Exit(1)
        }
        key := parts[0]
        value := parts[1]

        // 値が空文字の場合はダブルクォートで囲む
        if value == "" {
            envYamlContent.WriteString(fmt.Sprintf("%s: \"\"\n", key))
        } else if _, err := strconv.Atoi(value); err == nil {
            // 数値の場合はダブルクォートで囲む
            envYamlContent.WriteString(fmt.Sprintf("%s: \"%s\"\n", key, value))
        } else {
            envYamlContent.WriteString(fmt.Sprintf("%s: %s\n", key, value))
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    // YAML ファイルを出力
    err = os.WriteFile(yamlFile, []byte(envYamlContent.String()), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        os.Exit(1)
    }
    fmt.Printf("YAML file created at: %s\n", yamlFile)
}