package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

/*
3# rw{hW -> 8
3# r -> metade

*/

func finalStep(data string) string {
	mid := len(data) / 2
	crypt := data[:mid]
	
	// fmt.Println(crypt)
	// fmt.Println(mid)
	// fmt.Println(data)

	for i := mid; i < len(data); i++ {
		// fmt.Println(string(data[i]))
		position := rune(data[i]) - 1

		crypt += string(position)
	}

	// fmt.Println(crypt)

	return crypt
}

func secondStep(data string) string {
	crypt := ""
	for i := len(data) - 1; i >= 0; i-- {
		// fmt.Printf("%s", string(data[i]))
		crypt += string(data[i])
		// fmt.Println(i)
	}
	return crypt
}

func firstStep(data string) string {
	// TODO preciso separar o que é caracter (letras) do que não é

	crypt := ""
	notLetter := ""
	for i := 0; i < len(data); i++ {
		if unicode.IsLetter(rune(data[i])) {
			// fmt.Printf("%s", string(data[i]))
			// fmt.Printf("%d ", rune(data[i]))

			position := rune(data[i]) + 3

			// fmt.Printf("*%d\n", rune(data[i]))
			// fmt.Printf("%s\n", string(position))
			crypt += string(position)

		} else {
			notLetter += string(data[i])
		}
	}

	// fmt.Println(crypt)
	// fmt.Println(notLetter)

	crypt += notLetter

	// fmt.Println(crypt)

	return crypt

}

func promptIn() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

func rounds() int {
	var n int
	fmt.Scan(&n)
	return n
}

// func asciiTable() map[int]string {
// 	asciiMap := make(map[int]string)

// 	for i := 0; i < 127; i++ {
// 		asciiMap[i] = string(rune(i))
// 	}

// 	return asciiMap
// }

func main() {
	val := rounds()
	var c int
	
	for {
		if c == val+1 {
			break
		}
		
		s := firstStep(promptIn())
		s = secondStep(s)
		fmt.Println(finalStep(s))
		c++
	}
	
}
