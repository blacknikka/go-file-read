package main

import (
	"fmt"
	"os"

	"github.com/blacknikka/go-file-read/src/read"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You should indicate a directory path!")
		fmt.Println("	For instance,")
		fmt.Println()
		fmt.Println("		app.exe /path/to/directory")
		fmt.Println()
		os.Exit(1)
	}

	// ディレクトリのパスを指定して実行
	read.Read(os.Args[1])
}
