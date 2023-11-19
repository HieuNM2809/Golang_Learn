package main
import ("fmt")

func main() {
	i := 1

	// While 
	for {
		if i%2 == 0 {
			if i < 4 {
				i++
				continue
			}
			fmt.Println(i)
		}
		i++
		if i > 10 {
			break
		}
	}


	// foreach 
	numbers := [3]int{1, 2, 3}

    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
