package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
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

    file, err := os.Open(envFile)
    if err != nil {
        fmt.Println("Error reading .env file:", err)
        os.Exit(1)
    }
    defer file.Close()

    var envYamlContent strings.Builder
    scanner := bufio.NewScanner(file)

    // 元ファイルと同じ順序で行を出力したいため、1行ずつ読み込んで処理
    for scanner.Scan() {
        line := scanner.Text()

        unmarshaled, err := godotenv.Unmarshal(line)
        if err != nil {
            fmt.Println("Error parsing line:", err)
            os.Exit(1)
        }

        // プロパティが0個の場合はスキップ
        if len(unmarshaled) == 0 {
            continue
        }

        // 1行ずつ読んでいるためプロパティが2個以上になるはずが無いが、もし2個以上あればエラー
        if len(unmarshaled) > 1 {
            fmt.Println("Error parsing line: multiple properties found")
            os.Exit(1)
        }

        // プロパティが1個の場合は YAML 形式に変換
        for key, value := range unmarshaled {
            // ダブルクォートで囲む
            // 改行文字は\nに変換
            envYamlContent.WriteString(fmt.Sprintf("%s: \"%s\"\n", key, strings.ReplaceAll(value, "\n", "\\n")))
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    // YAML ファイルを出力
    yamlFile := envFile + ".yml"
    err = os.WriteFile(yamlFile, []byte(envYamlContent.String()), 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        os.Exit(1)
    }
    fmt.Printf("YAML file created at: %s\n", yamlFile)
}