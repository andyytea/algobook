package sort_go

func InsertionSortN(arr []int) []int {
    var i, j int
    for i = range arr {
       for j = i; j > 0 && arr[j] > arr[j-1]; j-- {
           arr[j], arr[j-1] = arr[j-1], arr[j]
       }
   }
   return arr
}
