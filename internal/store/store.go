package store

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
