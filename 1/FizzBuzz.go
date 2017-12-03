package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 引数の数を確認
	if len(os.Args) != 2 {
		// 画面に文字を出力
		fmt.Println("引数で一つだけ数字を入力してください。")
		fmt.Println(os.Args)
		// プログラムを終了する
		os.Exit(1)
	}
	// 引数で指定された数値を文字列から数値に変換
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// 画面に文字を出力
		fmt.Println("引数は整数で指定してください")
	}
	for i := 1; i < max+1; i++ {
		if i%3 == 0 || i%5 == 0 {
			// 3で割り切れ、5でも割り切れるとき
			// 画面にFizzBuzzを出力
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			// 3で割り切れるとき
			// 画面にFizzを出力
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			// 5で割り切れるとき
			// 画面にBuzzを出力
			fmt.Println("Buzz")
		} else {
			// 上記以外のとき
			// 画面にiの値を出力
			fmt.Println(i)
		}
	}
}
