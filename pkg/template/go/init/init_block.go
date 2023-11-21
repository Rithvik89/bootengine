package go_template

import (
	"bootengine/pkg/store"
	"fmt"
	"go/format"
	"os"
)

func InitWrite(path string) {
	_, err := os.Create(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	s3 := store.StringifyS3()

	fcode := fmt.Sprintf("package template %s", s3)
	formattedCode, err := format.Source([]byte(fcode))

	os.WriteFile(path, formattedCode, 0644)

}
