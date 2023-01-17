package types

func NewGenesisState(params Params, lastPairID uint64, pairs []Pair) *GenesisState {
	return &GenesisState{Params: params, LastPairId: lastPairID, Pairs: pairs}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams(), 0, nil)
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (m GenesisState) Validate() error {
	if err := m.Params.Validate(); err != nil {
		return err
	}
	return nil
}
