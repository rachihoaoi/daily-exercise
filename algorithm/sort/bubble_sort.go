package sort

func BubbleSort(array []int64) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func CookTailSort(array []int64) {
	begin, end := 0, len(array)-1
	for begin < end {
		newBegin, newEnd := begin, end

		for i := begin; i < end; i++ {
			if array[i] > array[i+1] {
				newEnd = i
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
		if newBegin == newEnd {
			break
		}

		end = newEnd

	}
}
