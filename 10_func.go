
// ----------------------- Hàm basic
// package main

// import "fmt"

// func main() {
//     a := 10
//     b := 20
//     fmt.Println(findMax(a, b))

//     c := 30
//     d := 40
//     fmt.Println(findMax(c, d))
// }

// func findMax(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }


// ----------------------- Hàm nâng cao 
// package main

// import "fmt"

// func main() {
//     a := []int{1, 2, 3, 4}
//     sum := findSum(a)
//     fmt.Println(sum)
// }

// func findSum(a []int) (sum int) {
//     for _, v := range a {
//         sum += v
//     }
//     return sum
// }

// ------  Rest parameter
// package main

// import "fmt"

// func main() {
//     sum := findSum("Hello", 1, 2, 3, 4, 5)
//     fmt.Println(sum)
// }

// func findSum(a string, numbers ...int) (sum int) {
//     for _, v := range numbers {
//         sum += v
//     }
//     fmt.Println(a)
//     return sum
// }

// ---------------- struc func
package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

