package main
import (
	"fmt"
)
func main(){
	// Khai báo một array key: string, value; int
	studentNameAgeMap := make(map[string] int)
	studentNameAgeMap  = map[string] int{
		"Peter" : 18,
		"David" : 20,
	}
	// gán lại data 
	studentNameAgeMap["Peter"] = 22

	//add key 
	studentNameAgeMap["Peter1"] = 18

	//delete key
	delete(studentNameAgeMap, "Peter1")

	//check isset ( destructuring )
	_, isExist := studentNameAgeMap["Peter1"]

	fmt.Println("List Array: ", studentNameAgeMap)

	fmt.Println("Key Peter is exist: ", isExist, _)
}