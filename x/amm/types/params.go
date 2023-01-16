package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramstypes.ParamSet = (*Params)(nil)

var (
	KeyFeeRate     = []byte("FeeRate")
	DefaultFeeRate = sdk.NewDecWithPrec(3, 3)
)

func NewParams(feeRate sdk.Dec) Params {
	return Params{FeeRate: feeRate}
}

func DefaultParams() Params {
	return NewParams(DefaultFeeRate)
}

func (params *Params) Validate() error {
	if err := validateFeeRate(params.FeeRate); err != nil {
		return err
	}
	return nil
}

func (params *Params) String() string {
	out, _ := yaml.Marshal(params)
	return string(out)
}

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyFeeRate, &params.FeeRate, validateFeeRate),
	}
}

func validateFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type : %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("fee rate must not be egative: %s", v)
	}
	return nil
}
