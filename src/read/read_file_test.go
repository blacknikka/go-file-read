package read

import (
	"os"
	"testing"
)

func TestAnalyzeAllDirectory_正常系(t *testing.T) {
	dir, _ := os.Getwd()

	err := AnalyzeAllDirectory(dir)
	if err != nil {
		t.Errorf("err発生")
	}
}
