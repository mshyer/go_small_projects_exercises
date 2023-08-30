package main

import (
	"fmt"

)

type yumCar interface {
	String() string
}

type bananaCar struct {
	num int16
	name string
}

func (b bananaCar) String() string {
	fmt.Println("firing")
	return b.name
}

func main() {


	banana := bananaCar{
		num: 33,
		name: "bananaCar",
	}

	var value interface{}

	value = banana
	fmt.Printf("%T\n", banana)

	switch str := value.(type) {
		case string:
			fmt.Println(str) 
		case yumCar:
			fmt.Println(str.String()) 
	}
}
