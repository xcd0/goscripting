/// 2>/dev/null ; gorun "$0" "$@" ; exit $?
// ↑shebangg

// 引数処理のテストを書きたい

package main

import (
	"fmt"
	"os"

	_ "github.com/fatih/color"            // CLIの出力をカラフルにする
	_ "github.com/jimlawless/whereami"    // 使用されているファイル名、行番号、関数などの情報をキャプチャする。エラーメッセージの改善に使用。
	_ "github.com/schollz/progressbar/v3" // 時間の掛かるオペレーション時に示するプログレスバー
	_ "github.com/spf13/cobra"            // CLIアプリケーションの入力処理、オプション、関連ドキュメントなどの複雑な処理の実装を簡単にする

	"a.a/goscripting/lib"
)

func main() {

	fmt.Println("Hello World!")

	cmd := lib.ParseArgs("ls -al")
	fmt.Println(cmd)

	if err := lib.RunCommand(cmd); err != nil {
		fmt.Println(err)
	}

	os.Exit(42)
}
