package main

import (
	"fmt"
	"strings"
)
func main() {
	fmt.Println("Welcome");
	result := caesar_cypher("L fdph, L vdz, L frqtxhuhg", 3)
	fmt.Println(result)
}

func caesar_cypher(text string, hop uint8) string {
  decoded_text := strings.Map(func(r rune) rune {
		if r == ',' || r == ' ' {
			return r
		}
		return r-rune(hop)
	}, text)
	return decoded_text;
}
