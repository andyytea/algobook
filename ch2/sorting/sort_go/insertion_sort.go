package sort_go

func InsertionSort1(arr []int) {
    var i, j int
    for i = range arr {
       for j = i; j > 0 && arr[j] > arr[j-1]; j-- {
           arr[j], arr[j-1] = arr[j-1], arr[j]
       }
   }
}

func InsertionSort2(arr []int) {
    var i, j int
    alen := len(arr)
    for i = 0; i < alen; i++ {
        for j = i; j > 0 && arr[j] > arr[j-1]; j-- {
            arr[j], arr[j-1] = arr[j-1], arr[j]
        }
    }
}
