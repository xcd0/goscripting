# golangをインタープリタ言語風に実行するメモ書き

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


## gomacro, gore, gorun
* [gomacro](https://github.com/cosmos72/gomacro)  
gomacroは `gomacro FILENAME.go` `#!/usr/bin/env gomacro` のようなシェバングに対応してる。
ただし標準のGo実装ではないため問題が発生することはあるらしい。  
gomacroはライセンスがMPLでGPLに近いが、gomacroを組み込まず実行ファイルを使用するのみのため問題ない。  
参考:https://www.infoq.com/jp/news/2020/06/go-scripting-language/  
という触れ込みだったが、実際に動かしてみるとHello World!が出力されなかった。  
とりあえず保留。

* [gore](https://github.com/x-motemen/gore)というgolangのREPL実装もあるが、
こちらはシェバングに対応していないように思われる。

* [gorun](https://github.com/erning/gorun)
とりあえず使ってみた感じこれが一番簡単に動いた。gorunはライセンスがGPLだが、gorunを組み込まず実行ファイルを使用するのみのため問題ない。  

gorunであればファイル先頭行に `/// 2>/dev/null ; gorun "$0" "$@" ; exit $?` を書くことで、
`chmod +x main.go; ./main.go`のように実行できた。
shebangとしてはわかりずらいが使えれば問題ない。

## 結論

1. `go install github.com/erning/gorun@latest` を実行。
2. 即時実行したいgolangのソースコードの先頭行に`/// 2>/dev/null ; gorun "$0" "$@" ; exit $?`を記載。
3. 権限を付与して呼び出し。

## おまけ

最新版のgolangのインストールをさっと実行したいとき
```
wget https://raw.githubusercontent.com/xcd0/go_compiler_install/master/install_lastest.sh && ./install_lastest.sh
```
でできるスクリプトを書いた。https://github.com/xcd0/go_compiler_install
