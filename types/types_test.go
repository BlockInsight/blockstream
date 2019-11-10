package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {

	var test map[string]Blockchain
	err := json.Unmarshal([]byte(`
	{
		"string":"ETH"
	}
	`), &test)

	require.NoError(t, err)
}
