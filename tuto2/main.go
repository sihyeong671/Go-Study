package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result32 float32 = floatNum32 + float32(intNum32)
	fmt.Println(result32)
	// 서로다른 숫자 타입을 더하는 것을 불가능하고 타입 캐스팅을 통해서 가능하게 한다.
	// tab을 통해 자동완성하면 vscode에서 자동으로 type cast를 시켜준다.

	var myString1 string = "Hello\nWorld"
	fmt.Println(myString1)
	var myString2 string = `Hello
World`
	// 백틱으로 나타내는 경우 안에 있는 문자 형식 그대로 출력 (enter 같은것)
	fmt.Println(myString2)

	fmt.Println(len("test"))                 // len 함수는 byte의 수를 출력
	fmt.Println(len("γ"))                    // 2
	fmt.Println(utf8.RuneCountInString("γ")) // 1
	// 문자열 길이를 구할 때 아스키코드를 확인해야함, 특수문자 주의!

	var myRune rune = 'a'
	fmt.Println(myRune) // 아스키코드 출력

	var myBool bool = true // true or false
	fmt.Println(myBool)

	// 기본값
	// int, 모든 숫자타입 = 0
	// string = ''
	// bool = false

	myVar := "text"
	fmt.Println(myVar)
	// var키워드 사용 필요없이 타입 추론 한다.

	v1, v2 := 1, 2
	fmt.Println(v1, v2)
	// 다중 입력 가능

	const myConst string = "const value"
	fmt.Println(myConst)
}
