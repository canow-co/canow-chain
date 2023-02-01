package helpers

import (
	"github.com/canow-co/canow-chain/app"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var (
	Codec    codec.Codec
	Registry types.InterfaceRegistry
)

func init() {
	encodingConfig := app.MakeEncodingConfig()
	Codec = encodingConfig.Marshaler
	Registry = encodingConfig.InterfaceRegistry
}
