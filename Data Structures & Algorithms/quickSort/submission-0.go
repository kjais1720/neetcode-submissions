// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

type Solution struct {

}

func NewSolution() *Solution {
	return &Solution{}
}

func Divide(arr []Pair, s, e int) {
	if e-s +1 <= 1 {
		return
	}

	left := s 

	for i := s; i < e ; i++ {
		if arr[i].Key < arr[e].Key {
			temp := Pair{
				Key: arr[i].Key,
				Value: arr[i].Value,
			}
			arr[i] = arr[left]
			arr[left] = temp
			left++
		}
	}
	temp := Pair{
				Key: arr[e].Key,
				Value: arr[e].Value,
			}
	arr[e] = arr[left]
	arr[left] = temp 

	Divide(arr, s, left-1)
	Divide(arr, left+1, e)
}

func QuickSort(pairs []Pair) []Pair {
	Divide(pairs, 0, len(pairs)-1)
	return pairs
}
