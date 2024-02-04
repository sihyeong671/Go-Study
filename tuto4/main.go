package main

import (
	"fmt"
	"time"
)

func main() {
	// 같은 타입의 고정 길이
	var intArr [3]int32
	// intArr := [...]int{1, 2, 3}
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	// Array는 연속 메모리에 저장됨
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(len(intSlice), cap(intSlice)) // 3 3
	intSlice = append(intSlice, 7)
	fmt.Println(len(intSlice), cap(intSlice)) // 4 6

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// var intSlice3 []int32 = make([]int32, 3, 8)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"] // key가 존재하면 ok = True
	// delete(myMap2, "Adam")
	fmt.Println(age, ok)

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Println(i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
