package store

import (
	"fmt"
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
		{Name: "python", Version: [2]int{3, 13}},
		{Name: "rust", Version: [2]int{1, 91}},
	}
	store := Store{
		Path: "/tmp/test-store",
		deps: make(map[string]Dep),
	}
	_ = store.Add(deps)
	has, _ := store.HasDeps(deps)
	testutil.Assert(t, has, "Expected store to contain deps")
}

func TestStoreSaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/store.json"

	store := NewStore(path)
	deps := []Dep{
		{Name: "python", Version: [2]int{3, 13}},
		{Name: "rust", Version: [2]int{1, 91}},
	}
	_ = store.Add(deps)

	err := store.Save()
	testutil.AssertNoErr(t, err, fmt.Sprintf("Failed to save"))

	loaded, err := LoadStore(path)
	testutil.AssertNoErr(t, err, "Expected to load")

	has, missing := loaded.HasDeps(deps)
	testutil.Assert(t, has, fmt.Sprintf("missing deps after load: %+v", missing))
}
