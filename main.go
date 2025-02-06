package main

import (
	"fmt"

	"github.com/saswatax/dsa-go/stack"
)

func main() {
	result := stack.StockSpan([]int{13, 15, 12, 14, 16, 8, 6, 4, 10, 30})
	fmt.Println(result)

}
