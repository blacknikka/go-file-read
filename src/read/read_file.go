package read

import (
	"fmt"
	"os"
	"path/filepath"
)

// Read ファイルを読み込む
func Read(dirPath string) {
	fmt.Println("Start reading a file.")
	fmt.Println("Target directory is " + dirPath)

	// ファイルをOpenする
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() == false {
			fmt.Println("A file is found : " + path)
			return nil
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
