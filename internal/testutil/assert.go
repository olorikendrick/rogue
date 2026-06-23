package testutil

import "testing"

func Assert(t *testing.T, condition bool, msg string) {
	t.Helper()
	if !condition {
		t.Fatalf("assertion failed: %s", msg)
	}
}

func AssertEq[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func AssertErr(t *testing.T, err error, msg string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error but got nil: %s", msg)
	}
}

func AssertNoErr(t *testing.T, err error, msg string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: unexpected error: %v", msg, err)
	}
}
