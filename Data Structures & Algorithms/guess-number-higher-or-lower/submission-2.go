/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l := 1
	r := n +1
	for l <= r {
		mid := (l+r)/2
		g := guess(mid)
		if g == -1 {
			r = mid
		} else if g == 1 {
			l = mid
		} else {
			return mid
		}
	}
	return -1
}
