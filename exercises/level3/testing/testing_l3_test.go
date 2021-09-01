package main

import (
	"testing"
)

func filterByAgeRange(fromAge, toAge int, ages []int) []int {
	response := make([]int, 0)
	for _, ageToCheck := range ages {
		if fromAge <= ageToCheck && ageToCheck <= toAge {
			response = append(response, ageToCheck)
		}
	}
	return response
}

// You can use the following function to Deep compare arrays
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestShouldPassByReturnArrayFiltered(t *testing.T) {
	ans := filterByAgeRange(5, 20, []int{2, 3, 5, 7, 11, 13})
	expected := []int{5, 7, 11, 13}
	if !Equal(ans, expected) {
		t.Errorf("Arrays not Equal ===>>> Expected: %v \tReceived: %v\n", expected, ans)
	}
}

func TestShouldFail(t *testing.T) {
	ans := filterByAgeRange(5, 12, []int{2, 3, 5, 7, 11, 13})
	expected := []int{2, 5, 7, 11, 13}
	if !Equal(ans, expected) {
		t.Errorf("Arrays not Equal ===>>> Expected: %v \tReceived: %v\n", expected, ans)
	}
}
