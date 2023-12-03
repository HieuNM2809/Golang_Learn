
// Goroutine ( async, await) : chạy bất đồng bộ
//1. go count("sheep") - chạy nền 

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	go count("sheep")
// 	go count("fish")
// 	time.Sleep(time.Second * 5)
// }

// func count(name string) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(name, i)
// 		time.Sleep(time.Second)
// 	}
// }

//--------------------------------------------------------------
// 2.  WaitGroup = Promise.all // đợi tất cả các job chạy xong 
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)  // thêm 2 job đợi

// 	go func() {
// 		count("sheep")
// 		wg.Done()  // hoàn thành => wg -1 
// 	}()

// 	go func() {
// 		count("fish")
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	fmt.Println("Done")
// }

// func count(name string) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(name, i)
// 		time.Sleep(time.Second)
// 	}
// }


//--------------------------------------------------------------
// 3. RLock , Lock // để controll các task vụ chạy đồng bộ
//Lock được sử dụng để đảm bảo rằng chỉ có một goroutine có thể truy cập vào critical section (phần mã quan trọng) tại một thời điểm.
//RLock cung cấp khả năng đọc đồng thời cho nhiều goroutine mà không bị cản trở.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter = 0
var m sync.RWMutex

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
