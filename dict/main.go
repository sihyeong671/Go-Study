package main

import (
	"fmt"

	"github.com/dict/main/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "greeting")

	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
