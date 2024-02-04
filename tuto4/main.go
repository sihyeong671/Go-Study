package main

import "fmt"

func main() {
	var myString string = "résumé"
	var indexed = myString[0]
	fmt.Println(indexed)
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}
	// String
	// ASCII - 7bits

	// UTF-8
	// 실제 바이트의 길이를 재기 때문에 rune 타입 사용하는 게 좋음
	var myRune = []rune("résumé")
	var runeIndexed = myRune[0]

	fmt.Println(runeIndexed)
	fmt.Printf("%v, %T\n", runeIndexed, runeIndexed)
	for i, v := range myRune {
		fmt.Println(i, v)
	}

}
