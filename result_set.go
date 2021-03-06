package floc

import "fmt"

// ResultSet is the set of possible results. This set is the simple
// implementation of Set with no check for duplicate values and it covers only
// basic needs of floc.
type ResultSet []Result

// NewResultSet constructs the set with given results. The function validates
// all result values first and panics on any invalid result.
func NewResultSet(results ...Result) ResultSet {
	// Validate results
	for _, res := range results {
		if !res.IsValid() {
			panic(fmt.Errorf("invalid result %s in result set", res.String()))
		}
	}

	return results
}

// Contains tests if the set contains the result.
func (set ResultSet) Contains(result Result) bool {
	for _, res := range set {
		if res == result {
			return true
		}
	}

	return false
}
