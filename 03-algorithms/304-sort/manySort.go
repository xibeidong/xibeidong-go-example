package sort

//选择排序
func selectionSort(arr []int) {
	if arr == nil || len(arr) == 1 {
		return
	}
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i //最小值的索引
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

//冒泡排序
func bubbleSort(arr []int) {
	if arr == nil || len(arr) == 1 {
		return
	}
	n := len(arr)
	for end := n - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

// 插入排序，对于局部有序的数组，效率更高，不是严格的O(n^2)
func insertSort(arr []int) {
	if arr == nil || len(arr) == 1 {
		return
	}
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

//归并排序一，递归法
func mergeSort(arr []int, L, R int) {
	if R == L {
		return
	}
	M := L + (R-L)>>1
	mergeSort(arr, L, M)
	mergeSort(arr, M+1, R)
	mergeArr(arr, L, M, R)
}

func mergeArr(arr []int, L, M, R int) {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1
	for p1 <= M && p2 <= R {
		if arr[p1] < arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= M {
		help[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}
	i = 0
	for j := L; j < R; j++ {
		arr[j] = help[i]
		i++
	}
}
