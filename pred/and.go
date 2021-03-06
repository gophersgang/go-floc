package pred

import floc "github.com/workanator/go-floc"

// And returns a predicate which chains multiple predicates into a contidion
// with AND logics. The result predicate finishes calculation of
// the condition as fast as the result is known. The function panics if
// the number of predicates is less than 2.
//
// The result predicate tests the condition as follows.
//   [PRED_1] AND ... AND [PRED_N]
func And(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		// More than 2 predicates
		return func(state floc.State) bool {
			for _, predicate := range predicates {
				if !predicate(state) {
					return false
				}
			}

			return true
		}
	} else if count == 2 {
		// 2 predicates
		return func(state floc.State) bool {
			return predicates[0](state) && predicates[1](state)
		}
	}

	// Require at least 2 predicates
	panic("And requires at least 2 predicates")
}
