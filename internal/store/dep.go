package store
import(
	"fmt"
)


type Dep struct {
	Name    string `json:"name"`
	Version [2]int `json:"version"`
}

func (d Dep) Key() string {
	    return fmt.Sprintf("%s@%d.%d", d.Name, d.Version[0], d.Version[1])
	    
}
