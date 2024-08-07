package main

import (
	"fmt"
	"math"
)

// user 计算最大用户数的函数
func user(R int, Ni []int, D int) int {
	idx := 0
	maxUser := 0
	mapping := make(map[int]int)
	n := R + 1

	// 将2的幂次映射到对应的用户数
	for idx < n {
		mapping[int(math.Pow(2, float64(idx)))] = Ni[idx]
		idx++
	}

	// 计算能直接满足条件的用户数，并将其从映射中移除
	for key, value := range mapping {
		if key >= D {
			maxUser += value
			mapping[key] = 0
		}
	}

	var x int64 = 0
	// 计算剩余的用户数总值
	for key, value := range mapping {
		x += int64(key) * int64(value)
	}

	// 计算剩余用户数能够组成的最大完整组
	maxUser += int(x / int64(D))
	return maxUser
}

func main() {
	var R int
	fmt.Print("Enter R: ")
	fmt.Scan(&R)

	Ni := make([]int, R+1)
	fmt.Printf("Enter %d values for Ni:\n", R+1)
	for i := 0; i <= R; i++ {
		fmt.Scan(&Ni[i])
	}

	var D int
	fmt.Print("Enter D: ")
	fmt.Scan(&D)

	maxUser := user(R, Ni, D)
	fmt.Println(maxUser)
}
