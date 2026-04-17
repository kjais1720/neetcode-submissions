func replaceElements(arr []int) []int {
    n := len(arr)
    ans := make([]int, n)
    rightMax := -1
    for i := n - 1; i >= 0; i-- {
        ans[i] = rightMax
        if arr[i] > rightMax {
            rightMax = arr[i]
        }
    }
    return ans
}