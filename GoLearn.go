// GoLearn
package main

import (
	"fmt"
)

const (
	A = iota
	B
	C
	d
)

type Bird struct {
}

func (b *Bird) Fly() {
	fmt.Println("Bird fly")
}

type IFly interface {
	Fly()
}

func main() {

	var ifly IFly = new(Bird)
	ifly.Fly()
	h := "2我们"
	fmt.Printf("%d %c\n", len(h), h[0])
	for i, ch := range h {
		fmt.Printf("%d %c\n", i, ch)
	}

	var a [5]int
	a = [5]int{1, 2, 3, 4}
	var b []int = a[:]
	for i, d := range b {
		fmt.Printf("%d %d\n", i, d)
	}

}
