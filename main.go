/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

package main

/*
# golangをスクリプト言語風に使用する仕組みの仮実装

bashやpythonなどスクリプト言語では型の恩恵を受けがたい。
強い型付けによりコンパイル時のエラーチェックがあり、これにより実行時エラーを削減できる。
強い型付け言語のスクリプト言語でTypescriptがあるが、
2021年現在のTypescriptの主な実装はMicrosoftによるJavascriptへトランスコンパイラであるため、
Javascriptの理解が必要で、また実行されるのはJavascriptであるため型システムによる実行速度の高速化の恩恵が受けられていない。

コンパイラ言語であれば強い型付け言語はいくつもある。
メジャーなところではC,C++,Java,Rustなどがある。

golangはコンパイラ言語であるがコンパイルが非常に高速であるため、
もともと`go run`コマンドにより即時実行が可能である。
しかしこれはあくまでコンパイルして実行しているため、on the flyな実行ではなく、
shebangが使えないなど、完全にスクリプト言語として使うにはいくつか問題がある。

shebangが使えればshellで
```sh
$ chmox +x main.go
$ ./main.go
```
のようにテキストファイルを直接実行できる用になるメリットがある。
`go run` コマンドでshebangを実行しようと、1行目に `#! /usr/bin/env go run` を記載すると
golangは#によるコメントアウトが機能しないこと、shebangは引数をとれないことが問題になる。
後者は`go run`コマンドに対するエイリアスを設定することで解決できる。
しかし前者はgolangの標準コンパイラでは解決が難しい。
`go run main.go`を実行するshellscriptを置いたり、makefileに書いたりといった方法もあるがスマートではない。

* [gomacro](https://github.com/cosmos72/gomacro)いう、
golangのインタープリタ実装を用いてインタープリタ言語としてgolangを使ってみる。
gomacroはライセンスがMPLでGPLに近いが、gomacroを使用するのみで組み込まないため問題ないと判断した。
gomacroは `gomacro FILENAME.go` `#!/usr/bin/env gomacro` のようなシェバングに対応してる。
ただし標準のGo実装ではないため問題が発生することはあるよう。
参考:https://www.infoq.com/jp/news/2020/06/go-scripting-language/

という事だったが、実際に動かしてみるとHello World!が出力されなかった。

* [gore](https://github.com/x-motemen/gore)というgolangのREPL実装もあるが、
こちらはシェバングに対応していないように思われる。

* [gorun](https://github.com/erning/gorun)
とりあえず使ってみた感じこれが一番簡単に動いた。

gorunであればファイル先頭行に `/// 2>/dev/null ; gorun "$0" "$@" ; exit $?` を書くことで、
`chmox +x main.go; ./main.go`のように実行できた。

*/

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
	os.Exit(42)
}
