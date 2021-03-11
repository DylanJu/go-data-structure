package heap

type Heap struct {
	nodes []int
}

func (heap *Heap) Insert(value int) {
	heap.nodes = append(heap.nodes, value)

	// find last index
	lastIndex := len(heap.nodes) - 1

	// swap
	heap.maxHeapifyUp(lastIndex)
}

func (heap *Heap) maxHeapifyUp(targetIndex int) {
	// find parent index of target index
	parentIndex := (targetIndex - 1) / 2

	// swap if target value is greater than parent value
	if heap.nodes[targetIndex] > heap.nodes[parentIndex] {
		heap.nodes[targetIndex], heap.nodes[parentIndex] = heap.nodes[parentIndex], heap.nodes[targetIndex]
		// recursive
		heap.maxHeapifyUp(parentIndex)
	}
}

func (heap *Heap) Pop() int {
	max := heap.nodes[0]

	lastIndex := len(heap.nodes) - 1
	heap.nodes[0] = heap.nodes[lastIndex]
	heap.remove(lastIndex)
	heap.maxHeapifyDown(0)

	return max
}

func (heap *Heap) maxHeapifyDown(parentIndex int) {
	lastIndex := len(heap.nodes) - 1
	leftIndex := (parentIndex * 2) + 1
	rightIndex := (parentIndex * 2) + 2

	if lastIndex < leftIndex && lastIndex < rightIndex {
		return
	}

	// left와 right 둘다 있을 경우
	// 둘 중에 큰걸 비교해서
	// 큰걸 parent에 assign 하고
	// 재귀호출
	if heap.nodes[leftIndex] > heap.nodes[rightIndex] {
		heap.nodes[parentIndex], heap.nodes[leftIndex] = heap.nodes[leftIndex], heap.nodes[parentIndex]
		heap.maxHeapifyDown(leftIndex)
	} else {
		heap.nodes[parentIndex], heap.nodes[rightIndex] = heap.nodes[rightIndex], heap.nodes[parentIndex]
		heap.maxHeapifyDown(rightIndex)
	}
}

func (heap *Heap) remove(index int) {
	heap.nodes = append(heap.nodes[:index], heap.nodes[index+1:len(heap.nodes)]...)
}

/*
	## 실수했던 Pop() ##

	아래처럼 할 경우 트리가 균형잡히지 않는다.
	pop/sort 이후의 트리 균형이 깨지는 것을 고려하지 않고 코딩되어 있다.
	heap.nodes = [63 16 50 14 8 48 20 9 1 5 7 34] 일 경우
	wrong pop을 할경우

*/

func (heap *Heap) wrongPop() int {
	max := heap.nodes[0]
	heap.wrongMaxHeapifyDown(0)

	return max
}

func (heap *Heap) wrongMaxHeapifyDown(parentIndex int) {
	lastIndex := len(heap.nodes) - 1
	leftIndex := (parentIndex * 2) + 1
	rightIndex := (parentIndex * 2) + 2

	if lastIndex < leftIndex && lastIndex < rightIndex {
		// parent 삭제
		heap.remove(parentIndex)
		return
	} else if lastIndex < leftIndex {
		// right를 parent로 assign 하고
		heap.nodes[parentIndex] = heap.nodes[rightIndex]
		// right 삭제
		heap.remove(rightIndex)
		return
	} else if lastIndex < rightIndex {
		// left를 parent로 assign 하고
		heap.nodes[parentIndex] = heap.nodes[leftIndex]
		// left 삭제
		heap.remove(leftIndex)
		return
	}

	// left와 right 둘다 있을 경우
	// 둘 중에 큰걸 비교해서
	// 큰걸 parent에 assign 하고
	// 재귀호출
	if heap.nodes[leftIndex] > heap.nodes[rightIndex] {
		heap.nodes[parentIndex] = heap.nodes[leftIndex]
		heap.wrongMaxHeapifyDown(leftIndex)
	} else {
		heap.nodes[parentIndex] = heap.nodes[rightIndex]
		heap.wrongMaxHeapifyDown(rightIndex)
	}
}
