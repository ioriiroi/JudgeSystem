package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

// プログラムの実行
func action(input_file string) (string, string) {
	// テストするPythonのファイルがあるディレクトリ名
	var testcode_address string = "testcode/"

	if len(os.Args) > 2 { // 引数がある場合はディレクトリ名を変更
		testcode_address = os.Args[2] + "/"
	}

	cmd := exec.Command("python3", testcode_address+"main.py") // 実行するコマンド

	var timelimit int = 2    // 実行時間制限
	var status string = "NA" // 実行時の状態
	var is_TLE bool = false  //TLEの判定

	var filename string = input_file

	// ファイルの中身の取得
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("error")
	}

	defer file.Close()

	cmd.Stdin = file //標準入力にファイルの中身を格納

	stdout, _ := cmd.StdoutPipe()
	buf := new(bytes.Buffer)

	go func() { // 2秒以上経ったらプログラム終了
		time.Sleep(time.Second * time.Duration(timelimit))
		is_TLE = true
		cmd.Process.Kill()
	}()

	// プログラム実行
	if err := cmd.Start(); err != nil {
		status = color.RedString("RE")
	}

	buf.ReadFrom(stdout) // 答えの書き込み

	// プログラム終了
	if err := cmd.Wait(); err != nil {
		status = color.RedString("RE")
	}

	out := buf.String() // 答えを文字列に変換

	if is_TLE {
		status = color.YellowString("TLE")
	}

	return out, status
}
