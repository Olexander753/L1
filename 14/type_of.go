package main

import "fmt"

// typeOf - define what type v is
func typeOf(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("%s is String type(%T)\n", v, v)
	case int:
		fmt.Printf("%d is Integer type(%T)\n", v, v)
	case bool:
		fmt.Printf("%v is Boolean type(%T)\n", v, v)
	case chan int:
		fmt.Printf("%v is Channel int type(%T)\n", v, v)
	}
}

func main() {
	str := "Hello"
	num := 42
	boolean := true
	ch := make(chan int)

	typeOf(str)
	typeOf(num)
	typeOf(boolean)
	typeOf(ch)
}
