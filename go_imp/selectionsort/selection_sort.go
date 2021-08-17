package selectionsort

// SelectionSort sorts the slice in ascending order
func SelectionSort(s []int) []int {
	var min int
	max := len(s)
	for i := 0; i < max; i++ {
		min = i
		for j := i + 1; j < max; j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
	return s
}
