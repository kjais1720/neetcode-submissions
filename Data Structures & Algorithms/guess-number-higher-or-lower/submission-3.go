/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 1, n
    for {
        m1 := l + (r-l)/3
        m2 := r - (r-l)/3
        if guess(m1) == 0 {
            return m1
        }
        if guess(m2) == 0 {
            return m2
        }
        if guess(m1)+guess(m2) == 0 {
            l = m1 + 1
            r = m2 - 1
        } else if guess(m1) == -1 {
            r = m1 - 1
        } else {
            l = m2 + 1
        }
    }
}