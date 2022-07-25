package main

import "fmt"

func main(){
	A := 5.
	B := 4.
	a := 5
	b := 4
	var C float64
	var boo1, boo2, boo3 bool

	C = (A+B) / 5
	fmt.Println(C)
	
	boo1 = (a == b)
	fmt.Println(boo1)
	boo2 = (a != b)
	fmt.Println(boo2)
	boo3 = (a >= b)
	fmt.Println(boo3)
	boo := boo1 && boo2
	fmt.Println(boo)

	var aa, bb byte = 0x1, 0x3
	c := (aa & bb) << 5
	fmt.Println(c)

	a = 100
	a *= 10
	fmt.Println(a)
	a >>= 2
	fmt.Println(a)
	a |= 1
	fmt.Println(a)

	var k int = 10
	var p = &k
	fmt.Println(p, *p)
	*p = 100
	fmt.Println(k)
}