package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}

	lo, hi := 0, len(a)-1

	for i, _ := range a {
		if a[i] < a[hi] {
			a[lo], a[i] = a[i], a[lo]
			lo++
		}
	}
	//[ < pivot ] <lo ( >= pivot) > [ >= pivot ] <hi == pivot>

	a[lo], a[hi] = a[hi], a[lo]
	//[ < pivot ] < a[lo] : pivot > [ >= pivot]

	quickSort(a[:lo])   // pivot左边
	quickSort(a[lo+1:]) // pivot右边
}

func main() {
	a := make([]int, 30)

	// 根据当前时间生成rand的种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// 赋随机值
	for i := 0; i < 30; i++ {
		a[i] = r1.Intn(10000)
	}

	quickSort(a)
	fmt.Println(a)
}
