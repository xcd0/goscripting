# golangをインタープリタ言語風に使用する仕組みの仮実装

bashやpythonなどインタープリタ言語では型の恩恵を受けがたい。
強い型付けによりコンパイル時のエラーチェックがあり、これにより実行時エラーを削減できる。
強い型付け言語のインタープリタ言語でTypescriptがあるが、
2021年現在のTypescriptの主な実装はMicrosoftによるJavascriptへトランスコンパイラであるためJavascriptの理解が必要である。

コンパイラ言語であれば強い型付け言語はいくつもあり、メジャーなところではC,C++,Java,Rustなどがある。

golangはコンパイラ言語であるがコンパイルが非常に高速であるため、`go run`コマンドにより即時実行が可能である。
しかしこれはあくまでコンパイルして実行しているため、完全にインタープリタ言語として使うにはいくつか問題がある。
また、`#`によるコメントが不可であるためshebangが使えない。
このため`chmod +x ./main.go; ./main.go`のような実行できない。
この問題を解決する手段を探した。

* [gomacro](https://github.com/cosmos72/gomacro)いう、
参考:https://www.infoq.com/jp/news/2020/06/go-scripting-language/
gomacroはライセンスがMPLでGPLに近いが、gomacroを組み込まず実行ファイルを使用するのみのため問題ない。
gomacroは `gomacro FILENAME.go` `#!/usr/bin/env gomacro` のようなシェバングに対応してる。
ただし標準のGo実装ではないため問題が発生することはあるらしい。
という触れ込みだったが、実際に動かしてみるとHello World!が出力されなかった。
とりあえず保留。

* [gore](https://github.com/x-motemen/gore)というgolangのREPL実装もあるが、
こちらはシェバングに対応していないように思われる。

* [gorun](https://github.com/erning/gorun)
とりあえず使ってみた感じこれが一番簡単に動いた。

gorunであればファイル先頭行に `/// 2>/dev/null ; gorun "$0" "$@" ; exit $?` を書くことで、
`chmox +x main.go; ./main.go`のように実行できた。
shebangとしてはわかりずらいが使えれば問題ない。


