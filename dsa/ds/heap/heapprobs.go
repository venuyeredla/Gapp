package heap

import "fmt"

func HeapSort(arr []int) {
	minHeap := NewMinHeap(len(arr))
	for _, ele := range arr {
		minHeap.Push(ele, nil)
	}
	for i := 0; i < len(arr); i++ {
		entry := minHeap.PopPriority()
		fmt.Printf("%v ,", entry)
		//minHeap.fixDown(0)
	}
}
