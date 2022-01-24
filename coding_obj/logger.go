package coding_obj

import (
	"fmt"
	"os"
	"time"
)

// Go言語では、 yyyy とかではなく、定められた数をそこに置くのらしい☆（＾～＾）
const timeStampLayout = "2006-01-02 15:04:05"

// write - ログファイルに書き込みます。
func write(filePath string, text string, args ...interface{}) {
	// TODO ファイルの開閉回数を減らせないものか。
	// 追加書込み。
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// tはtime.Time型
	t := time.Now()

	s := fmt.Sprintf(text, args...)
	s = fmt.Sprintf("[%s] %s", t.Format(timeStampLayout), s)
	fmt.Fprint(file, s)
	defer file.Close()
}
