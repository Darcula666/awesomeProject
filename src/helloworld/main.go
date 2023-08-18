// 从控制台读取输入:
package main

import (
	"fmt"
	"github.com/Darcula666/calculator"
)

func main() {
	firstName := "John"
	updateName(&firstName)
	fmt.Println(firstName)
	total := calculator.Sum(3, 5)
	fmt.Println(total)
}

func updateName(name *string) {
	*name = "David"
}
