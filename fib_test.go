package fib

import (
	"testing"
)

func TestFibCiclo(t *testing.T) {
	assertions := []struct {
		assert int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
	}
	for _, assertion := range assertions {
		execution := FibClico(assertion.assert)
		if execution != assertion.expect {
			t.Errorf("Error se esperaba %d y tenemos %d \n ", execution, assertion.assert)
		}
	}
}
