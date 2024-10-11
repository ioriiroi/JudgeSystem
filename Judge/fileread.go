package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ファイルを開き、その内容を文字列の配列で返す。
func fileread(name string) string {
	var file *os.File

	file, err := os.Open(name)

	if err != nil {
		fmt.Println("error")
	}

	defer file.Close()

	buf := new(bytes.Buffer)

	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println("error")
	}

	out := buf.String()

	return out
}
