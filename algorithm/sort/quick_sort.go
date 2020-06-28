package sort

func QuickSort(array []int, begin, end int) {
	if begin < end {
		index := getIndex(array, begin, end)
		QuickSort(array, begin, index-1)
		QuickSort(array, index+1, end)
	}
}

func getIndex(array []int, begin, end int) int {
	for begin < end {
		for begin < end && array[begin] <= array[end] {
			end--
		}
		array[begin], array[end] = array[end], array[begin]
		for begin < end && array[begin] <= array[end] {
			begin++
		}
		array[begin], array[end] = array[end], array[begin]
	}
	return begin
}
