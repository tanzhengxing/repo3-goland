package main

import "fmt"

var arr = [...]int{1, 5, 3, 4, 7, 77, 89, 77, 78, 45}

func main() {
	//1、冒泡排序
	//bubbleSort()
	//bubbleSort_rev(&arr)

	//2、选择排序
	//selectSort()
	//selectSort_rev(&arr)
	fmt.Print("haaaaaa")

	//3、插入排序
	//insertSort()
	//insertSort_rev(&arr)

	//4、希尔排序
	//shellSort()
	//shellSort_rev(&arr)

	//5、归并排序
	//MergeSort(&arr, 0, 9)
	//fmt.Println(arr)
	//mergeSort_rev(&arr, 0, len(arr)-1)

	//6、快速排序
	//quickSort(&arr, 0, 9)
	//fmt.Println(arr)
	//quickSort_rev(&arr,0,9)

	//堆排序
	heapSort(&arr)
	//fmt.Println(arr)
	//heapSort_rev(&arr)
	fmt.Println(arr)

}
func heapSort_rev(arr *[10]int) {
	//先初始化最大堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap_rev(arr, i, len(arr)-1)
	}

	//重排所有的堆
	for j := len(arr) - 1; j > 0; j-- {
		arr[j], arr[0] = arr[0], arr[j]
		adjustHeap_rev(arr, 0, j-1)
	}
}
func adjustHeap_rev(arr *[10]int, low int, high int) {
	parent := low
	temp := arr[parent]
	child := parent*2 + 1
	for child <= high {
		if child < high && arr[child] < arr[child+1] {
			child++
		}
		if child <= high && arr[child] > temp {
			arr[parent] = arr[child]
			parent = child
			child = parent*2 + 1
		} else {
			break
		}
	}
	arr[parent] = temp
}
func quickSort_rev(arr *[10]int, left int, right int) {
	if left < right {
		var standard int = findStandard(arr, left, right)
		quickSort_rev(arr, left, standard-1)
		quickSort_rev(arr, standard+1, right)
	}
}
func findStandard(arr *[10]int, left int, right int) int {
	privot := arr[left]
	if left < right {
		for left < right && arr[right] >= privot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
			left++
		}
		for left < right && arr[left] <= privot {
			left++
		}
		if left < right {
			arr[right] = arr[left]
			right++
		}
	}
	arr[left] = privot
	return left
}
func mergeSort_rev(arr *[10]int, left int, right int) {
	if left < right {
		middle := (left + right) / 2
		mergeSort_rev(arr, left, middle)
		mergeSort_rev(arr, middle+1, right)
		mergeSort(arr, left, middle, right)
	}
}
func mergeSort(arr *[10]int, left int, middle int, right int) {
	var s [10]int
	var k int
	i := left
	j := middle + 1
	for i <= middle && j <= right {
		if arr[i] < arr[j] {
			s[k] = arr[i]
			i++
		} else {
			s[k] = arr[j]
			j++
		}
		k++
	}

	for i <= middle {
		s[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		s[k] = arr[j]
		j++
		k++
	}
	for i := 0; i < k; i++ {
		arr[left+i] = s[i]
	}

}
func shellSort_rev(arr *[10]int) {
	num := len(arr) / 2
	for num >= 1 {
		for i := num; i < len(arr); i += num {
			index := i - num
			curVal := arr[index+num]
			for index >= 0 && arr[index] > curVal {
				arr[index+num] = arr[index]
				index -= num
			}
			arr[index+num] = curVal
		}
		num = num / 2
	}
}
func insertSort_rev(arr *[10]int) {
	for i := 1; i < len(arr); i++ {
		index := i - 1
		curVal := arr[index+1]
		for index >= 0 && arr[index] > curVal {
			arr[index+1] = arr[index]
			index--
		}
		arr[index+1] = curVal
	}
}
func selectSort_rev(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
func bubbleSort_rev(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

/*func heapSort(arr *[10]int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr)-1)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjustHeap(arr, 0, i-1)
	}
}
func adjustHeap(arr *[10]int, low int, high int) {
	parent := low
	child := parent*2 + 1
	temp := arr[parent]
	for child < high {
		if child < high && arr[child] < arr[child+1] {
			child++
		}
		if child < high && arr[child] > temp {
			arr[parent] = arr[child]
			parent = child
			child = parent*2 + 1
		} else {
			break
		}
	}
	arr[parent] = temp

}*/
func heapSort(arr *[10]int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr)-1)
	}
	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(arr, 0, j-1)
	}
}

func adjustHeap(arr *[10]int, low int, high int) {
	parent := low
	child := parent*2 + 1
	temp := arr[parent]
	for child < high {
		if child < high && arr[child] < arr[child+1] {
			child = child + 1
		}
		if child < high && arr[child] > temp {
			arr[parent] = arr[child]
			parent = child
			child = parent*2 + 1
		} else {
			break
		}
	}
	arr[parent] = temp

}

//快速排序
/*func quickSort(arr *[10]int, left int, right int) {
	if left < right {
		var standard int = getStandard(arr, left, right)
		quickSort(arr, left, standard-1)
		quickSort(arr, standard+1, right)
	}
}
func getStandard(arr *[10]int, left int, right int) int {
	privot := arr[left]
	for left < right {
		for left < right && arr[right] >= privot {
			right--
		}
		if left < right {
			arr[left] = arr[right]
		}

		for left < right && arr[left] <= privot {
			left++
		}
		if left < right {
			arr[right] = arr[left]
		}
	}
	arr[left] = privot
	return left

}*/

func quickSort(arr *[10]int, left int, right int) {
	if left < right {
		pindex := getStandard(arr, left, right)
		quickSort(arr, left, pindex-1)
		quickSort(arr, pindex+1, right)
	}
}
func getStandard(arr *[10]int, left int, right int) int {
	privot := arr[left]
	for left < right {
		for left < right && privot < arr[right] {
			right--
		}
		if left < right {
			arr[left] = arr[right]
			left++
		}
		for left < right && privot > arr[left] {
			left++
		}
		if left < right {
			arr[right] = arr[left]
			right--
		}
	}
	arr[left] = privot
	return left
}

//归并排序
/*func MergeSort(arr *[10]int, left int, right int) {
	if left < right {
		middle := (left + right) / 2
		MergeSort(arr, left, middle)
		MergeSort(arr, middle+1, right)
		merge(arr, left, middle, right)
	}
}
func merge(arr *[10]int, left int, middle int, right int) {
	if left >= right {
		return
	}
	var s *[10]int = &[10]int{}
	var k int
	i := left
	j := middle + 1
	for i <= middle && j <= right {
		if arr[i] < arr[j] {
			s[k] = arr[i]
			i++
		} else {
			s[k] = arr[j]
			j++
		}
		k++
	}
	for i <= middle {
		s[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		s[k] = arr[j]
		j++
		k++
	}
	for i := 0; i < k; i++ {
		arr[left] = s[i]
		left++
	}
}*/

func MergeSort(arr *[10]int, left int, right int) {
	if left < right {
		middle := (left + right) / 2
		MergeSort(arr, left, middle)
		MergeSort(arr, middle+1, right)
		Merge(arr, left, middle, right)
	}
}
func Merge(arr *[10]int, left int, middle int, right int) {
	var s [10]int
	var k int
	i := left
	j := middle + 1
	for i <= middle && j <= right {
		if arr[i] < arr[j] {
			s[k] = arr[i]
			i++
		} else {
			s[k] = arr[j]
			j++
		}
		k++
	}
	for i <= middle {
		s[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		s[k] = arr[j]
		j++
		k++
	}
	for i := 0; i < k; i++ {
		arr[left+i] = s[i]
	}
}

//希尔排序
func shellSort() {
	num := len(arr) / 2
	for num >= 1 {
		for i := num; i < len(arr); i += num {
			index := i - num
			currval := arr[index+num]
			for index >= 0 && arr[index] > currval {
				arr[index+num] = arr[index]
				index -= num
			}
			arr[index+num] = currval
		}
		num /= 2
	}
}

//插入排序
func insertSort() {
	for i := 1; i < len(arr); i++ {
		index := i - 1
		currval := arr[i]
		for index >= 0 && currval < arr[index] {
			arr[index+1] = arr[index]
			index--
		}
		arr[index+1] = currval
	}
}

//选择排序
func selectSort() {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

//冒泡排序
func bubbleSort() {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
