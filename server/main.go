package main

import (
	"fmt"
	"regexp"
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

}

func map_practice() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)
}

func main() {
	// fmt.Println("Hello, World!", time.Now())
	// fmt.Println(user.Current())

	// fmt.Println(i, f64, s, t, f)

	// foo()
	// string_practice()
	map_practice()

}
