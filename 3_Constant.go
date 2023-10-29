/*
	1.Định nghĩa hằng số (constant)
	2. Khai báo hằng số
	3. Đặc điểm của hằng số
	4. Kiểu dữ liệu Enum
	5. Tự dộng khởi tạo giá trị cho Enum bằng iota
*/
package main
import ( 
	"fmt"
)
const (
	_ = iota  // _ : khai báo bỏ qua biến đầu
	Red      // Viết hoa biến thì cho phép ngoài scope có thể truy cập được
	blue
	yellow
	orange
)
// const (
// 	Red     = 1 
// 	blue    = 2
// 	yellow  = 3
// 	orange  = 4
// )
func main() {
	fmt.Printf("%v, %T\n", Red, Red)
	fmt.Printf("%v, %T\n", blue, blue)
	fmt.Printf("%v, %T\n", yellow, yellow)
	fmt.Printf("%v, %T\n", orange, orange)
}


