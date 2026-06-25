
package store 

type Manifest struct {
    Packages map[string]Package `toml:"packages"`
}

type Package struct {
    Name     string   `toml:"name"`
    Versions []string `toml:"versions"`
    Recipe   Recipe   `toml:"recipe"`
}

type Recipe struct {
    URL        string            `toml:"url"`
    BuildTools map[string][]string `toml:"build_tools"` 
    Build      map[string]Build  `toml:"build"`        
}

type Build struct {
    Commands []string `toml:"commands"`
}
