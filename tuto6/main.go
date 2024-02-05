package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfor owner
	owner
	// 이렇게 바로 struct를 넣으면 내부 객체에 바로 접근하도록 만들 수 있음
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

// struct gasEngine의 메서드 만드는 방법
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

// 다른 타입에 동일한 메서드가 있는 경우 interface 사용할 필요 있음
type engine interface {
	milesLeft() uint8
}

// e에 특성 타입이 아닌 인터페이스를 넣어 해당 메서드를 불러오게 만들 수 있음
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func main() {
	// 순서 맞춰서 arg방식으로 써줄 수도 있고
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, owner: owner{"Alex"}}
	// 따로 지정해서 값 할당할 수 있음
	myEngine.mpg = 10

	// 아래처럼 할당 가능하지만 재사용이 불가능함
	var tmpEngine = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println(tmpEngine.mpg, tmpEngine.gallons)
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Printf("Total miles left in tank : %v\n", myEngine.milesLeft())

	var myEletricEngine electricEngine = electricEngine{mpkwh: 25, kwh: 15}
	canMakeIt(myEletricEngine, 50)
}
