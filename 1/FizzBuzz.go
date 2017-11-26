package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 引数の数を確認
	if len(os.Args) != 2 {
		fmt.Println("引数で一つだけ数字を入力してください。")
		fmt.Println(os.Args)
		os.Exit(1)
	}
	// 引数で指定された数値を文字列から数値に変換
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("引数は整数で指定してください")
	}
	for i := 1; i < max+1; i++ {
		// 3で割り切れ、5でも割り切れるとき
		if i%3 == 0 || i%5 == 0 {
			fmt.Println("FizzBuzz")
			// 3で割り切れるとき
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			// 5で割り切れるとき
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			// 上記以外のとき
		} else {
			fmt.Println(i)
		}
	}
}
