package types

func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{Params: params}
}

func DefaultGenesis() *GenesisState {
	return NewGenesisState(DefaultParams())
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}
