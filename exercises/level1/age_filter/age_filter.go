package ageFilter

/*
 Function that will filter a Slice of ages that are between the range.
 The function will receive two numbers and a slice of ages as parameters.
 It should return the ages between the range
*/

func AgeFilter(n1, n2 int, ages []int) (string, []int) {
	newAgeSlice := []int{}

	if n1 > n2 {
		return "n1 should be greater than n2", newAgeSlice
	}

	for i := 0; i < len(ages); i++ {
		if n1 <= ages[i] && ages[i] <= n2 {
			newAgeSlice = append(newAgeSlice, ages[i])
		}
	}

	return "Ages: ", newAgeSlice
}
