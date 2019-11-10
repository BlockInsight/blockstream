package types

import (
	"encoding/json"

	"github.com/libs4go/errors"
)

var errVendor = errors.WithVendor("blockstream-types")

// Errors .
var (
	ErrBlockchain = errors.New("invalid blockchain name", errVendor, errors.WithCode(-1))
)

// MarshalJSON .
func (blockchain Blockchain) MarshalJSON() ([]byte, error) {

	return json.Marshal(blockchain.String())
}

// UnmarshalJSON .
func (blockchain *Blockchain) UnmarshalJSON(b []byte) error {
	println("?????")
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	val, ok := Blockchain_value[s]

	if !ok {
		return errors.Wrap(ErrBlockchain, "invalid blockchain %s", s)
	}

	*blockchain = Blockchain(val)

	return nil
}
