package main

import "fmt"

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("Create panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("Chia cho mot so la (")
}

// panic để xuất 1 lỗi, 
//còn defer để gọi lại sau và nhận lỗi bằng recover 
