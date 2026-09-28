// Package invariant provides assertions for programmer errors.
package invariant

// Failure is the value panicked by a failed assertion.
type Failure struct {
	// Msg describes the failed assertion.
	Msg string
}

// Error returns the assertion failure message.
func (f Failure) Error() string {
	return "assertion failed: " + f.Msg
}
