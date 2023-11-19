package main
import (
	"fmt"
)
// tao 1 struct
type Student struct {
    name     string
	age      int
	sex      bool
	subject  []string
}
func main(){
	// Cách 1: Khởi tạo trực tiếp khi khai báo
	studentA := Student {
		name    : "Hieu Minh",
		age     : 18,
		sex     : true,
		subject : []string { "Math",  "English"},
	}
	fmt.Println("studentA: ", studentA)

	// Cách 2: Khởi tạo rỗng và gán giá trị sau
	studentB         := Student {}
	studentB.name     = "Anna"
	studentB.age      = 20
	studentB.sex      = false
	studentB.subject  =  []string { "Math",  "English"}

	fmt.Println("studentB: ", studentB)
}