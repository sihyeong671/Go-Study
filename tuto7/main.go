package main

import "fmt"

func main() {
	// nil pointer will be error
	var p *int32 = new(int32) // default = 0
	var i int32

	fmt.Printf("The value p points to is: %v\n", *p) // dereferencing
	fmt.Printf("The value if i is: %v\n", i)
	p = &i
	*p = 10

	var k int32 = 2
	k = i // copy
	fmt.Println(k)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // copy pointer
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// function에 넘기는 파라미터는 기본적으로 call by value로 복사를 한다. 포인터를 써주면 효율적으로 코딩 가능
}
