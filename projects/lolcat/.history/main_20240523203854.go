package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"syreclabs.com/go/faker"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune) {
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
	fmt.Println()
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gorainbow")
	}

	for i := 0; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	output := strings.Join(phrases[:], "; ")

	fmt.Println()
}
