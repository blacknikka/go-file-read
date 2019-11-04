package read

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Read ファイルを読み込む
func Read() {
	fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	f, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	fmt.Println(string(b))
}
