package main
import ( 
	"fmt"
)
func main() {
	//c1
	// points := [5]int{1,2,3,4,5}
	//c2
	// points := [...]int{1,2,3,4,5}

	// fmt.Printf("%v, %T\n", points, points)
	// fmt.Printf("item 1 : %v, %T\n", points[1], points[1], )
	// fmt.Println(len(points))  // chiều dài 

    //c3
	//make: cấp phát dữ liệu
	a := make([]int, 0)
	fmt.Printf("a %v, %v, %v \n", a, len(a), cap(a))
	a = append(a, 1) // Thêm 1 phần tử
	fmt.Printf("a %v, %v, %v \n", a, len(a), cap(a))
	a = append(a, []int{2, 3, 4, 5, 6}...)  // Thêm nhiều phần tử
	fmt.Printf("a %v, %v, %v \n", a, len(a), cap(a))
}