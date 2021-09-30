package lib

import "strings"

// golangで`ls -al` のように外部コマンドを呼び出す場合
// out, err := exec.Command("ls", "-la").Output()
// のように[]stringに分割する必要がある。
// この関数で "ls -al" を "ls", "-al" のように分割する
func ParseArgs(str string) []string {
	// 一度スペース区切りした後引用符ああれば結合する
	spaced := strings.Fields(str)
	appendOut := func(in *[]string, out *[]string) {
		var bytesbuf = make([]byte, 0, 100) // アロケーション対策に100byteのバッファを確保
		for _, v := range *in {
			bytesbuf = append(bytesbuf, v...)
			bytesbuf = append(bytesbuf, ' ') // 間にスペース入れる
		}
		*out = append(*out, string(bytesbuf))
	}

	var out, buf []string
	dq, sq := false, false
	for _, s := range spaced {
		appended := false
		f, b := s[0], s[len(s)-1] // 最初と最後の文字を保持
		if f == '"' {
			dq = true
			buf = append(buf, s)
			appended = true
		} else if f == '\'' {
			sq = true
			buf = append(buf, s)
			appended = true
		}
		if b == '"' {
			if dq {
				dq = false
			}
			if !appended { // このsの1文字目が"だったらすでにappendされている appendされていないとき"の終わりを表す
				buf = append(buf, s)
				appendOut(&buf, &out)
			}
		} else if b == '\'' {
			if sq {
				sq = false
			}
			buf = append(buf, s)
			sq = false
			if !appended { // このsの1文字目が'だったらすでにappendされている appendされていないとき'の終わりを表す
				buf = append(buf, s)
				appendOut(&buf, &out)
			}
		}
	}
	appendOut(&buf, &out)

	return out
}
