package main

import (
	"fmt"
	"strings"
	// "github.com/gSpera/morse"
)

func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
	natoMap := make(map[string]string)
	output := ""
	for idx := 0; idx < len(nato); idx++ {
		natoMap[string(nato[idx][0])] = nato[idx]
	}
	for idx := 0; idx < len(words); idx++ {
		stringChar := string(words[idx])
		if natoMap[strings.ToUpper(stringChar)] != "" {
			output += natoMap[strings.ToUpper(stringChar)]
			output += " "
		} else {
			if stringChar != " " {
				output += stringChar
			}
		}
	}
	return output
}

func main() {
	// text := "MORSE IS AWESOME"
	// fmt.Println(morse.ToMorse(text))
	// DecodeMorse(".... . -.--   .--- ..- -.. .")
	// fmt.Println(morse.DefaultMorse)
	// fmt.Println(morse.DefaultMorse['P'])
	fmt.Println(ToNato("haha"))
	fmt.Println(ToNato("If you can read"))
}
