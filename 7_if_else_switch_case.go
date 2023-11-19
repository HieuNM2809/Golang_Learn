// package main
// import ("fmt")


// func main() {
// 	tamp := make(map[string] int)
// 	// Sử dụng ký tự dấu bằng (=) để gán giá trị và kiểm tra sự tồn tại của phần tử
// 	_, isExist := tamp["key1"]

// 	if isExist {
// 		fmt.Println("if true")
// 	} else {
// 		fmt.Println("if false")
// 	}
// 	fmt.Println("Hello world")
// }

// ------------------------------------------------------------------------------------

package main

import "fmt"

const targetNumber = 50 // Số mục tiêu để đoán

func main() {
	guess := getUserGuess()
	checkGuess(guess)
}

func getUserGuess() int {
	var guess int
	fmt.Println("Nhập số bạn đoán:")
	fmt.Scan(&guess)
	return guess
}

func checkGuess(guess int) {
	switch {
		case guess < 10 || guess > 100:
			fmt.Println("Con số bạn đoán nhỏ hơn 10 hoặc lớn hơn 100")
		case guess < targetNumber:
			fmt.Println("Con số mục tiêu lớn hơn con số bạn đoán")
		case guess > targetNumber:
			fmt.Println("Con số mục tiêu nhỏ hơn con số bạn đoán")
		default:
			fmt.Println("Bạn đã đoán đúng!")
	}
}
