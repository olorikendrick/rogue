package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type storeJSON struct {
	Path string         `json:"path"`
	Deps map[string]Dep `json:"deps"`
}

func (s *Store) UnmarshalJSON(data []byte) error {
	var j storeJSON
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	s.Path = j.Path
	s.deps = j.Deps
	return nil
}

type Store struct {
	Path string
	deps map[string]Dep
}

func NewStore(path string) *Store {
	return &Store{
		Path: path,
		deps: make(map[string]Dep),
	}
}

func (s *Store) HasDeps(deps []Dep) (bool, []Dep) {
	unknown := []Dep{}
	for _, dep := range deps {
		_, exists := s.deps[dep.Key()]
		if !exists {
			unknown = append(unknown, dep)
		}

	}
	if len(unknown) != 0 {
		return false, unknown
	} else {
		return true, nil
	}
}

func (s *Store) Add(deps []Dep) error {
	for _, dep := range deps {
		key := dep.Key()
		s.deps[key] = dep
	}
	return nil
}

func (s *Store) Save() error {
	j := storeJSON{
		Path: s.Path,
		Deps: s.deps,
	}
	data, err := json.Marshal(&j)
	if err != nil {
		return fmt.Errorf("Could not parse Store %+v:%w", j, err)
	}
	err = os.WriteFile(s.Path, data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write store %+v to file %s: %w", j, s.Path, err)
	}
	return nil

}

func LoadStore(path string) (*Store, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("Could not open file :%w", err)
	}

	var store Store
	err = store.UnmarshalJSON(data)
	if err != nil {
		return nil, fmt.Errorf("Could not parse store:%w", err)
	}

	return &store, nil
}
