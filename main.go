/// 2>/dev/null ; gorun "$0" "$@" ; exit $? // shebang

// 引数処理のテストを書きたい

package main

import (
	"fmt"
	"os"

	_ "github.com/fatih/color"            // CLIの出力をカラフルにする
	_ "github.com/jimlawless/whereami"    // 使用されているファイル名、行番号、関数などの情報をキャプチャする。エラーメッセージの改善に使用。
	_ "github.com/schollz/progressbar/v3" // 時間の掛かるオペレーション時に示するプログレスバー
	_ "github.com/spf13/cobra"            // CLIアプリケーションの入力処理、オプション、関連ドキュメントなどの複雑な処理の実装を簡単にする
)

func main() {

	fmt.Println("Hello World!")

	err := RunCommand()

	os.Exit(42)
}
