/*
	1. Định nghĩa biến trong ngôn ngữ lập trình
	2. Cú pháp khai báo biến
	3. Global and block scope
	4. Shadow
	5. Declared and not used
	6. Export global scope
	7. Naming convention
	8. Convert type
*/

//---------------------------------------------------------
// 1. Định nghĩa biến trong ngôn ngữ lập trình
// 2. Cú pháp khai báo biến

// Khai báo 1 biến 
//(1) var i int           -- Khai báo 
//(2) var i int = 10      -- Khai báo và khởi tạo
//(3) i := 10             -- Kháo báo và tự ép kiểu -> sử dùng trong 
//(4) var student2 = 111  -- Kháo báo và tự ép kiểu -> sử dùng trong và ngoài func



package main

import "fmt"

var j int = 1 

func main() {
	i := 1
	var j int = 2
	//j := 2
	var student2 = 111 //type is inferred
	fmt.Printf("value: %v, type: %T  \n", student2, student2)
	fmt.Println("Hello, 世界", i+1, j)
}

//-----------------------------------------------------
// 3. Global and block scope
// 4. Shadow  - Global được khai báo lại trong func

// package main

// import "fmt"

// var j int = 1 
// var (
// 	a int = 1
// 	b string = "1"
// 	c int = 1
// )

// func main() {
// 	fmt.Println("Hello, 世界", a,b,c,j)
// }

//----------------------------------------------------
// 6. Export global scope
// var Num int = 10
// -> Viết hoa chữ cái đầu đề các pakeage khác truy suất được

//----------------------------------------------------
// 7. Naming convention
// -> camel

//----------------------------------------------------
// 8. Convert type
// hoặc sử dụng strconv
// package main

// import "fmt"


// func main() {
// 	var a int  = 100
// 	var b float32  = float32(a)
// 	fmt.Printf("%v %T", b, b)
// }

