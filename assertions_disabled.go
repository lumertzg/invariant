//go:build invariant_disabled

package invariant

// Enabled reports whether invariant checks are enabled.
const Enabled = false

// True does nothing when invariant_disabled is set.
func True(bool, string) {}

// Unreachable does nothing when invariant_disabled is set.
func Unreachable(string) {}
