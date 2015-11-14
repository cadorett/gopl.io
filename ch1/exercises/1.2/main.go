// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Exercise 1.2: Modify the echo program to print the index
// and value of each of its arguments, one per line. 

// Modified by cado

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d, %s\n", i, os.Args[i])
	}
}

//!-
