/*
Exercise 1.1

Modify the echo program to also print os.Args[0], the name of the command that invoked it.

echo3: https://github.com/adonovan/gopl.io/blob/master/ch1/echo3/main.go
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
