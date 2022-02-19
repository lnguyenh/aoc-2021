package day15

type heapNode struct {
	name      string
	totalRisk int
}

type aocHeap []heapNode

func (h aocHeap) Len() int           { return len(h) }
func (h aocHeap) Less(i, j int) bool { return h[i].totalRisk < h[j].totalRisk }
func (h aocHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *aocHeap) Push(x interface{}) {
	*h = append(*h, x.(heapNode))
}

func (h *aocHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
