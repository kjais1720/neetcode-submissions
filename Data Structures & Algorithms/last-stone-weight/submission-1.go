
type Heap struct {
	arr []int
}

func (h *Heap) Len() int {
	return len(h.arr)
}

func (h *Heap) Push(n int) {
	h.arr = append(h.arr, n)
	j := h.Len() - 1
	for i := h.Len() / 2; i >= 1 && h.arr[i] < h.arr[j]; {
		temp := h.arr[i]
		h.arr[i] = h.arr[j]
		h.arr[j] = temp
		j = i
		i /= 2
	}
}

func (h *Heap) Pop() int {
	if h.Len() == 0 {
		return -1
	} else if h.Len() <= 2 {
		res := h.arr[1]
		h.arr = h.arr[:1]
		return res
	}
	res := h.arr[1]
	h.arr[1] = h.arr[h.Len()-1]
	h.arr = h.arr[:h.Len()-1]
	for i := 1; 2*i < h.Len(); {
		if h.arr[i] < h.arr[2*i] && (2*i+1 >= len(h.arr) || h.arr[2*i+1] <= h.arr[2*i]) {
			temp := h.arr[i]
			h.arr[i] = h.arr[2*i]
			h.arr[2*i] = temp
			i = 2 * i
		} else if 2*i+1 < len(h.arr) && h.arr[i] < h.arr[2*i+1] && h.arr[2*i+1] > h.arr[2*i] {
			temp := h.arr[i]
			h.arr[i] = h.arr[2*i+1]
			h.arr[2*i+1] = temp
			i = 2*i + 1
		} else {
			break
		}
	}

	return res
}

func Heapify(stones []int) Heap {
	if len(stones) == 0 {
		return Heap{
			arr: stones,
		}
	}

	stones = append(stones, stones[0])
	for i := len(stones) / 2; i >= 1; i-- {
		l := 2 * i
		j := i
		for l < len(stones) {
			if stones[j] < stones[l] && (l+1 >= len(stones) || stones[l+1] <= stones[l]) {
				temp := stones[j]
				stones[j] = stones[l]
				stones[l] = temp
				j = l
				l = 2 * l
			} else if l+1 < len(stones) && stones[j] < stones[l+1] && stones[l+1] > stones[l] {
				temp := stones[j]
				stones[j] = stones[l+1]
				stones[l+1] = temp
				j = l + 1
				l = 2*l + 1
			} else {
				break
			}
		}
	}

	return Heap{
		arr: stones,
	}
}

func lastStoneWeight(stones []int) int {
	hp := Heapify(stones)
	for hp.Len() > 2 {
		s1 := hp.Pop()
		s2 := hp.Pop()
		abs := math.Abs(float64(s2 - s1))
		if abs > 0 {
			hp.Push(int(abs))
		}
	}
	if hp.Len() < 2 {
		return 0
	}
	return hp.arr[1]
}