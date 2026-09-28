//go:build !invariant_disabled

package invariant

// Enabled reports whether invariant checks are enabled.
const Enabled = true

// True panics with Failure when cond is false.
func True(cond bool, msg string) {
	if !cond {
		panic(Failure{Msg: msg})
	}
}

// Unreachable panics with Failure because its call site must not be reached.
func Unreachable(msg string) {
	panic(Failure{Msg: msg})
}
