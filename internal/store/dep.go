package store



type Dep struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (d Dep) Key() string {
	return d.Name + "@" + d.Version
}
