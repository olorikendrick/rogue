package store

import (
	"testing"

	"github.com/olorikendrick/rogue/internal/testutil"
)

func TestStoreNew(t *testing.T) {
	store := NewStore("/tmp/test-store")
	testutil.Assert(t, store != nil, "expected store to not be nil")
	testutil.AssertEq(t, store.Path, "/tmp/test-store")
}

func TestStoreHasDeps(t *testing.T) {
	deps := []Dep{
		Dep{
			Name:    "python",
			Version: "3.11",
		},
		Dep{
			Name:    "rust",
			Version: "1.91",
		},
	}
	store := Store{
		Path: "/tmp/test-store",
		deps: make(map[string]Dep),
	}
	_ = store.Add(deps)
	has, _ := store.HasDeps(deps)

	testutil.Assert(t, has, "Expected store to contain deps ")
}
