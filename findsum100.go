package main

import (
	"fmt"
	"strconv"
)

func findSumSoFar(equation string, start int, end int, sum int, total int) {
	if sum == total {
		fmt.Println(equation + "=" + strconv.Itoa(total))
		return
	}
	if start == end {
		return
	}
	var numSoFar int = 0
	var sumSoFar int = 0
	for ; start < end; start++ {
		sumSoFar += start
		numSoFar = numSoFar * 10 + start
		findSumSoFar(equation + "+" + strconv.Itoa(numSoFar), start +1, end, sum + numSoFar, total)
		findSumSoFar(equation + "-" + strconv.Itoa(numSoFar), start +1, end, sum - numSoFar, total)
	}
}

func main() {
	findSumSoFar("", 1, 10, 0, 100)
}
