package main

import (
	"fmt"
	"os"
	"strconv"
)

type Saiyan struct {
	Name  string
	Power int
}

func Super(s *Saiyan) {
	fmt.Println("within", s)
	fmt.Println(*s)
	s = &Saiyan{"Gohan", 1000}
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}

	// for _, x := range "😊" {
	// 	fmt.Println("iterate")
	// 	fmt.Printf("%v\n", x)
	// }
	smiley := "😊"
	bytes := []byte(smiley)
	// runes := []rune(smiley)
	for idx := 0; idx < len(bytes); idx++ {
		fmt.Println(bytes[idx], "new line")
	}
	fmt.Println(bytes)
	// fmt.Println(string(240))
	fmt.Println(string(126))

	fmt.Println([]byte("é"))
	fmt.Println(byte('a'))
	fmt.Println(rune('😊'))
	fmt.Println(string([]byte{240, 159, 152, 138}))

	// myString := "abcdefg"
	// bytesBoi := []byte(myString)
	// for idx := len(bytesBoi) - 1; idx >= 0; idx-- {
	// 	fmt.Println(string(bytesBoi[idx]))
	// }

}
