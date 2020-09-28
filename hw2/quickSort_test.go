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
