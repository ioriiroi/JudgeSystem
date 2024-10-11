package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func judge() {
	// テストケースの入出力が置いてあるディレクトリ
	var testcase_address string = "testcase/"

	if len(os.Args) > 1 { // 引数がある場合はディレクトリ名を変更
		testcase_address = os.Args[1] + "/"
	}

	//ディレクトリのファイル名を取得
	files_in, _ := os.ReadDir(testcase_address + "in")

	files_out, _ := os.ReadDir(testcase_address + "out")

	var extension int = 4 //拡張子 [.txt] 分
	var txt string = ".txt"
	var address_in string = testcase_address + "in/"
	var address_out string = testcase_address + "out/"

	var correct_num int = 0
	var files_num int = len(files_in)

	for num := 0; num < files_num; num++ {
		input := address_in + files_in[num].Name() // 入力のファイル名読み込み

		fmt.Print(input[len(address_in):] + ": ") // 入力ファイル名の表示

		var my_answer, status = action(input)            // 実行する 答えと実行時の状態(TLEなど)をもらう
		var output = address_out + files_out[num].Name() // 出力のファイル名読み込み

		correct := fileread(output) // 答えの取得

		if input[len(input)-extension:] != txt { // txtファイル以外は読み込まない
			continue
		}

		//fmt.Println(input, my_answer)
		//fmt.Println(output, correct)

		if status == "NA" { // 正常に実行できたら答えを比較
			if my_answer == correct {
				status = color.GreenString("AC")
				correct_num++
			} else {
				status = color.RedString("WA")
			}
		}
		fmt.Println(status) // 結果の表示
	}
	fmt.Println("Accuracy:", correct_num, "/", files_num)
}

func main() {
	judge()
}
