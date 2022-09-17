package validate_subsequence


func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
    
    arrayIdx := 0
    sequenceIdx := 0

    for arrayIdx < len(array) && sequenceIdx < len(sequence) {

        if array[arrayIdx] == sequence[sequenceIdx] {
            sequenceIdx++
        }
        arrayIdx++
    }
    
	return sequenceIdx == len(sequence)
}

