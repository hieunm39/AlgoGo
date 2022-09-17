package validate_subsequence

import "testing"


func IsValidSubsequenceTest(t *testing.T) {

	array := []int{5, 1, 22, 25, 6, -1, 8}
	sequence := []int{1, 6, -1, 10}

	output := IsValidSubsequence(array, sequence) 

	expect := true

	if output != expect {
        t.Errorf("wrong expected")
    }
}