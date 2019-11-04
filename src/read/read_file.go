package read

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/blacknikka/go-file-read/src/analyze"
)

// AnalyzeAllDirectory ファイルを読み込む
func AnalyzeAllDirectory(dirPath string) (err error) {
	fmt.Println("Start reading a file.")
	fmt.Println("Target directory is " + dirPath)

	// ファイルをOpenする
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic("Error detected when a file/dir is opened : " + path)
		}

		if info.IsDir() == false && filepath.Ext(info.Name()) == ".txt" {
			fmt.Println("A file is found : " + path)

			// ファイルをOpenする
			fileContent, err := os.Open(path)

			if err != nil {
				fmt.Println("error")
				return err
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
		return err
	}

	return nil
}
