package read

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/blacknikka/go-file-read/src/analyze"
)

// AnalyzeAllDirectory ファイルを読み込む
func AnalyzeAllDirectory(dirPath string) {
	fmt.Println("Start reading a file.")
	fmt.Println("Target directory is " + dirPath)

	// ファイルをOpenする
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() == false && filepath.Ext(info.Name()) == ".txt" {
			fmt.Println("A file is found : " + path)

			// ファイルをOpenする
			fileContent, err := os.Open(path)

			if err != nil {
				fmt.Println("error")
			}
			defer fileContent.Close()

			byteArray, err := ioutil.ReadAll(fileContent)
			if err == nil {
				analyze.Analyze(string(byteArray))
			}

			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("error")
		return
	}
	// defer f.Close()

	// b, err := ioutil.ReadAll(f)
	// fmt.Println(string(b))
}
