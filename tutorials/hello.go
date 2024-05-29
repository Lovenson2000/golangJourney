package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {

	// Taking input
	reader := bufio.NewReader(os.Stdin)

	pl("What is your name ? ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name)

	pl("How old are you ? ")
	ageStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ageStr = strings.TrimSpace(ageStr)

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		log.Fatal(err)
	}
	pl("Hello", name, "\nYou are", age, "years old.")
}
