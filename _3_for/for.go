package main

import (
	"fmt"
	"math"
)

func main() {
	const k = 10
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		const d = 3e20 / k
		fmt.Println(d)
		//fmt.Println(int64(d))
		fmt.Println(math.Sin(k))
		break
	}


}
