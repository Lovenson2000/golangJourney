package main

import (
	"fmt"
	"strings"

	"github.com/enodata/faker"
)

func main() {
	var phrases []string

	for i := 0; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().phrases()...)
	}

	fmt.Println(strings.Join(phrases[:], "; "))
}
