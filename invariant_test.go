//go:build !invariant_disabled

package invariant

import "testing"

func TestTruePasses(t *testing.T) {
	if !Enabled {
		t.Fatal("Enabled = false, want true without invariant_disabled")
	}

	True(true, "must not panic")
}

func TestTruePanicsWithFailure(t *testing.T) {
	requireFailure(t, func() {
		True(false, "broken state")
	}, "broken state")
}

func TestUnreachablePanicsWithFailure(t *testing.T) {
	requireFailure(t, func() {
		Unreachable("unexpected branch")
	}, "unexpected branch")
}

func requireFailure(t *testing.T, call func(), wantMsg string) {
	t.Helper()

	defer func() {
		t.Helper()
		value := recover()
		failure, ok := value.(Failure)
		if !ok {
			t.Fatalf("panic = %T (%v), want invariant.Failure", value, value)
		}
		if failure.Msg != wantMsg {
			t.Errorf("Failure.Msg = %q, want %q", failure.Msg, wantMsg)
		}
		if got, want := failure.Error(), "assertion failed: "+wantMsg; got != want {
			t.Errorf("Failure.Error() = %q, want %q", got, want)
		}
	}()

	call()
	t.Fatal("call did not panic")
}
