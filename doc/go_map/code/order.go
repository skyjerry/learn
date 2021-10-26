package main

import "fmt"

func main() {
	m := map[string]string{
		"A": "A",
		"B": "B",
		"C": "C",
		"D": "D",
		"E": "E",
	}

	for _, v := range m {
		fmt.Println(v) // maybe [A,B,C,D,E] or [C,D,E,A,B]. It's not certain
	}
}
