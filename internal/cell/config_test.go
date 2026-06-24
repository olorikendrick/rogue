package cell

import (
	"testing"

	"github.com/olorikendrick/rogue/internal/testutil"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("./testData/config.json")
	testutil.AssertNoErr(t, err, "LoadConfig")

	testutil.AssertEq(t, config.ID, "agent-1")
	testutil.AssertEq(t, len(config.Deps), 2)
	testutil.AssertEq(t, config.Deps[0].Name, "python")
}

func TestValidate(t *testing.T) {
	config, err := LoadConfig("./testData/config.json")
	testutil.AssertNoErr(t, err, "LoadConfig")

	config.Timeout = 0
	testutil.AssertErr(t, config.Validate(), "Validate should fail when Timeout is 0")

	config.Timeout = 30
	config.ID = ""
	testutil.AssertErr(t, config.Validate(), "Validate should fail when ID is empty")
}
