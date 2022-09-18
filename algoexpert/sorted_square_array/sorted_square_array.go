package sortedsquarearray


func SortedSquaredArray(array []int) []int {

    result := make([]int, len(array))

    left := 0
    right := len(array) - 1

    for i := len(array) - 1; i >= 0; i-- {

        if abs(array[left]) > abs(array[right]) {
            result[i] = array[left] * array[left]
            left++
        } else {
            result[i] = array[right] * array[right]
            right--
        }
    }
    
	return result
}


func abs(a int) int {
    if a < 0 {
        a = a * -1
    }
    return a
}