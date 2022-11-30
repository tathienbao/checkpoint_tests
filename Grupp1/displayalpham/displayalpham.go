/*displayalpham
Instructions
Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

Usage
$ go run . | cat -e
aBcDeFgHiJkLmNoPqRsTuVwXyZ$
$*/

package main

import "github.com/01-edu/z01"

func main() {
	s := "aBcDeFgHiJkLmNoPqRsTuVwXyZ"
	for _, c := range s {
		z01.PrintRune(c)
	}
}
