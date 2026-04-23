func countStudents(students []int, sandwiches []int) int {

	i := 0
	for ; len(students) > 0 && len(sandwiches) > 0; {
		if students[0] == sandwiches[0] {
			students = reduceQueue(students)
			sandwiches = reduceQueue(sandwiches)
			i = 0
		} else {
			students = rotateQueue(students)
			i++
		}
		if i == len(students) {
			break
		}
	}
	return len(students)
}

func rotateQueue(q []int) []int{
	if len(q) <= 1 {
		return q
	}
	return append(q[1:len(q)], q[0])
}

func reduceQueue(q []int) []int {
	if len(q) == 1 {
		return []int{}
	}
	return q[1: len(q)]
}