package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {

	array := []int{49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22, 49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22, 49, 38, 65, 97, 23, 22, 76, 1, 5, 8, 2, 0, -1, 22, 123, 41, 2, 1, 6, 8, 4, 2, 5, 7, 1, 21, 532, 235, 7, 3, 2, 6}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)

}
