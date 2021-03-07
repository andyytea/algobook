package insertionsort

// InsertionSort sorts a list of integers in ascending order
func InsertionSort(l []int) []int {
	var j int
	for i := range l {
		j = i
		for (j > 0) && (l[j] < l[j-1]) {
			l[j], l[j-1] = l[j-1], l[j]
			j = j - 1
		}
	}
	return l
}
