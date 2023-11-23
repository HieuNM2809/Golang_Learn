package main

import "fmt"

func main() {
	///
	fmt.Println("VD 1: Defer để chạy câu lệnh sau")
	fmt.Println("Hello, 世界 1")
	defer fmt.Println("Hello, 世界 2.1")
	defer fmt.Println("Hello, 世界 2.2")
	defer fmt.Println("Hello, 世界 2.3")

	fmt.Println("Hello, 世界 3")


	///
	fmt.Println("VD 2: Defer nó sẽ lấy snapshot biến khi chạy đến dòng đó: ")
	a := 1
	defer fmt.Println(a)

	a = 10
	fmt.Println(a)

}

//defer            : chạy sau khi code chạy xong
//panic("message") : ngừng và xuất thông tin ( và die exit php)
