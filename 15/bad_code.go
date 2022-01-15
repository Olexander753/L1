package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

var justString string

// createHugeString - creates huge string from radom chars
func createHugeString(size int) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(generator.Intn('z'-'a'+1)))
	}
	return hugeString.String()
}

func SomeFunc() {
	v := createHugeString(1 << 10)

	// justString = v[:100]

	// trouble with non-UTF-8 characters so better conver to rune
	//justString = string([]rune(v)[:100])

	// trouble with memory size of underlying array
	justString = string(append([]rune{}, []rune(v)[:100]...))
	// tmp := make([]rune, 100)
	// copy(tmp, []rune(v)[:100])
	// justString = string(tmp)

	fmt.Println("justString = ", justString)
	fmt.Println("cap(justString) ", cap([]rune(justString)))
	fmt.Printf("%p\n", &justString)
	fmt.Println("Size in bytes ", int(unsafe.Sizeof(justString))+len(justString))
	fmt.Println()
	fmt.Println("v = ", v)
	fmt.Println("cap(v) ", cap([]rune(v)))
	fmt.Printf("%p\n", &v)
	fmt.Println("Size in bytes ", int(unsafe.Sizeof(v))+len(v))

}

func main() {
	SomeFunc()
}
