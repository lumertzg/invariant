//go:build invariant_disabled

package invariant

import "testing"

func TestDisabledInvariants(t *testing.T) {
	if Enabled {
		t.Fatal("Enabled = true, want false with invariant_disabled")
	}

	checks := 0
	messages := 0
	condition := func() bool {
		checks++
		return false
	}
	message := func() string {
		messages++
		return "does not panic"
	}

	True(condition(), message())
	Unreachable(message())

	if checks != 1 {
		t.Errorf("condition evaluations = %d, want 1", checks)
	}
	if messages != 2 {
		t.Errorf("message evaluations = %d, want 2", messages)
	}
}
