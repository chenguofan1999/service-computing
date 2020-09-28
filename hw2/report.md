# Golang实现快速排序

## 首先编写测试

```go
package main

import (
	"math/rand"
	"testing"
	"time"
)

//测试倒序排序
func TestSort1(t *testing.T) {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	quickSort(array)
	for i := 0; i < 10; i++ {
		if sortedArray[i] != array[i] {
			t.Error("unsorted!")
		}
	}
}

//测试特定数组排序
func TestSort2(t *testing.T) {
	array := []int{1, 9, 8, 2, 3, 7, 6, 4, 5}
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	quickSort(array)
	for i := 0; i < 9; i++ {
		if sortedArray[i] != array[i] {
			t.Error("unsorted!")
		}
	}
}

//测试随机数组排序
func TestSort3(t *testing.T) {
	// 构建一个大小为100的切片， 每个数都是1 - 10000间的随机数
	slice := make([]int, 100)

	// 根据当前时间生成rand的种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// 赋随机值
	for i := 0; i < 100; i++ {
		slice[i] = r1.Intn(10000)
	}

	// 排序
	quickSort(slice)

	// 检查有序
	for i := 0; i < 99; i++ {
		if slice[i] > slice[i+1] {
			t.Error("unsorted!")
		}

	}

}

```

这里设计了三个测试函数：
- TestSort1 : 测试特定**倒序**数组的排序
- TestSort2 : 测试特定**乱序**数组的排序
- TestSort3 : 测试**随机生成**的数组的排序

测试的对象是需要实现的quickSort函数。



## 假装根据测试文件实现quickSort函数

```go
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

```

除了给出quickSort的实现外还在main函数中写了个简单的测试，对随机生成的三十个整数的数组排序并输出。

## main函数的运行结果

```bash
chen@ChenSurface:~/service-computing/hw2$ go run quickSort.go
[90 208 702 882 1552 1659 1974 1998 2002 2248 2305 2563 3330 3515 4015 4389 5651 5726 5829 6073 6935 7247 7380 7609 7621 7777 8644 9037 9233 9878]
```

输出了有序的数组，不过这还不够证明算法的正确性，还需要进行完整的测试

## go test 结果

```bash
chen@ChenSurface:~/service-computing/hw2$ go test -v
=== RUN   TestSort1
--- PASS: TestSort1 (0.00s)
=== RUN   TestSort2
--- PASS: TestSort2 (0.00s)
=== RUN   TestSort3
--- PASS: TestSort3 (0.00s)
PASS
ok      quickSort       0.009s
```

通过了全部测试，以上就是一次成功的测试驱动式开发。