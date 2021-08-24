package sort_go

func InsertionSortRange(arr []int) {
    var i, j int
    for i = range arr {
       for j = i; j > 0 && arr[j] > arr[j-1]; j-- {
           arr[j], arr[j-1] = arr[j-1], arr[j]
       }
   }
}

func InsertionSortIndices(arr []int) {
    var i, j int
    alen := len(arr)
    for i = 0; i < alen; i++ {
        for j = i; j > 0 && arr[j] > arr[j-1]; j-- {
            arr[j], arr[j-1] = arr[j-1], arr[j]
        }
    }
}
