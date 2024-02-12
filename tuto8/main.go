// goroutine

package main

import (
	"fmt"
	"sync"
	"time"
)

// RWMutex사용가능
var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{}

var results = []string{}

func main() {

	for i := 0; i < 10; i++ {
		newID := fmt.Sprintf("id%d", i)
		dbData = append(dbData, newID)
	}

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal Execution is: %v\n", time.Since(t0))
	fmt.Printf("\nTotal results are %v\n", results)
	fmt.Println("\nTotal result len is", len(results))

}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock() // 전역변수 혹은 공유메모리 사용하는 경우에 같은 메모리에 접근시 덮어쓰는 경우가 생김
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}
