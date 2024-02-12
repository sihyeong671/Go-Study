package main

import (
	"fmt"
	"math/rand"
	"time"
)

// goroutine with channel
// hold data
// thread safe
// listen for data

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {

	// deadlock
	// var c = make(chan int)
	// c <- 1
	// var i = <-c
	// fmt.Println(i)

	// silly
	// var c = make(chan int, 5)
	// go process(c)
	// for i := range c {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second * 1)
	// }
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)

	}
	sendMessage(chickenChannel, tofuChannel)
}

// func process(c chan int) {
// 	defer close(c)
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// 	// close(c)
// }

func checkTofuPrice(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func checkChickenPrice(website string, chickenCahnnel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenCahnnel <- website
			break
		}
	}
}

func sendMessage(chickenCahnnel chan string, tofuChannel chan string) {
	select { // channel을 사용할때의 switch
	case website := <-chickenCahnnel:
		fmt.Printf("\nFound a deal on chicken at %s", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail Sent: Found deal on tofu at %v", website)
	}
}
