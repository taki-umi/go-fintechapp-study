package main

import (
	"fmt"
	"os/user"
	"time"
)

const (
	Pi       = 3.14
	Username = "test_user"
	Password = "test_pass"
)

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}

func string_practice() {
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	s = stirngs.Replace(s, "H", "X")
}

func main() {
	fmt.Println("Hello, World!", time.Now())
	fmt.Println(user.Current())

	fmt.Println(i, f64, s, t, f)

	foo()
	string_practice()
}
