package types

import (
	"github.com/libs4go/errors"
	"github.com/libs4go/scf4go"
)

// ChainIDValidator .
type ChainIDValidator map[Blockchain][]string

// GetChainIDValidator .
func GetChainIDValidator(config scf4go.Config) (ChainIDValidator, error) {
	validChainIDs := make(map[Blockchain][]string)

	var buff map[string][]string

	if err := config.Scan(&buff); err != nil {
		return nil, errors.Wrap(err, "get support chainIDs error")
	}

	if len(buff) == 0 {
		return nil, errors.Wrap(ErrChainIDs, "support chainIDs not config")
	}

	for name, val := range buff {
		blockchain, ok := Blockchain_value[name]

		if !ok {
			return nil, errors.Wrap(ErrBlockchain, "invalid blockchain %s", name)
		}

		validChainIDs[Blockchain(blockchain)] = val
	}

	return validChainIDs, nil
}

// CheckChainID .
func (validator ChainIDValidator) CheckChainID(blockchain Blockchain, chainID string) bool {
	list, ok := validator[blockchain]

	if !ok {
		return false
	}

	for _, valid := range list {
		if valid == chainID {
			return true
		}
	}

	return false
}
