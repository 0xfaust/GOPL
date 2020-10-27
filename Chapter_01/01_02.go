/*
Exercise 1.2

Modify the echo program to print the index and value of each of its arguments, one per line.

echo3: https://github.com/adonovan/gopl.io/blob/master/ch1/echo3/main.go
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
