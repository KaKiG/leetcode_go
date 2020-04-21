package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	buildings := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	fmt.Println(getSkyline(buildings))
}

// PLZ read the code carefully !!!!!!!!!!!!!!!!!
type entry struct {
	key   int
	value int
}

type hashHeap struct {
	data []entry
	hash map[int]int
}

func (h hashHeap) Len() int { return len(h.data) }
func (h hashHeap) Less(l, r int) bool {
	return h.data[l].key > h.data[r].key
}
func (h hashHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.hash[h.data[i].key] = i
	h.hash[h.data[j].key] = j
}

func (h *hashHeap) Push(x interface{}) {
	e := x.(entry)
	h.hash[e.key] = len(h.data)
	h.data = append(h.data, e)
}

func (h *hashHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *hashHeap) erase(key int) entry {
	idx := h.hash[key]
	rt := heap.Remove(h, idx)
	delete(h.hash, key)
	return rt.(entry)
}

func (h *hashHeap) get(key int) (int, bool) {
	idx, ok := h.hash[key]
	if ok {
		return h.data[idx].value, ok
	} else {
		return 0, ok
	}
}

func (h *hashHeap) update(key int, val int) {
	idx := h.hash[key]
	h.data[idx].value = val
}

func getSkyline(buildings [][]int) [][]int {
	points := make([][]int, 0)
	for _, v := range buildings {
		points = append(points, []int{v[0], -v[2]})
		points = append(points, []int{v[1], v[2]})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	hh := &hashHeap{}
	hh.hash = make(map[int]int)
	e0 := entry{0, 1}

	heap.Push(hh, e0)

	pre_h := 0
	cur_h := 0
	rt := make([][]int, 0)

	for _, v := range points {
		if v[1] < 0 {
			key := -v[1]
			value, ok := hh.get(key)
			if ok {
				hh.update(key, value+1)
			} else {
				heap.Push(hh, entry{key, 1})
			}
		} else {
			key := v[1]
			value, _ := hh.get(key)
			if (value - 1) == 0 {
				hh.erase(key)
			} else {
				hh.update(key, value-1)
			}
		}

		cur_h = hh.data[0].key
		//fmt.Println(cur_h)
		if cur_h != pre_h {
			rt = append(rt, []int{v[0], cur_h})
			pre_h = cur_h
		}
	}

	return rt
}
