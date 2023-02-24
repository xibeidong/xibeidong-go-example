package sort

import (
	"log"
	"testing"
)

var arr []int

func TestMain(m *testing.M) {
	arr = make([]int, 1024)
	for i := 0; i < 1024; i++ {
		arr[i] = i
	}
	m.Run()
}

func TestSelectionSort(t *testing.T) {
	arr := []int{23, 45, 11, -23, 3, 7, 5, 90, 12}
	selectionSort(arr)
	log.Println(arr)
}
func BenchmarkSelectionSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		selectionSort(arr)
	}
	//log.Println(arr)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{23, 45, 11, -23, 3, 7, 5, 90, 12}
	bubbleSort(arr)
	log.Println(arr)
}

func TestInsertSort(t *testing.T) {
	arr := []int{23, 45, 11, -23, 3, 7, 5, 90, 12}

	insertSort(arr)
	log.Println(arr)
}

func BenchmarkInsertSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		insertSort(arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{23, 45, 11, -23, 3, 7, 5, 90, 12}

	mergeSort(arr, 0, len(arr)-1)
	log.Println(arr)
}
