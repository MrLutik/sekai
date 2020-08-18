package types

import (
	"github.com/KiraCore/cosmos-sdk/codec"
	"github.com/KiraCore/cosmos-sdk/codec/types"
	sdk "github.com/KiraCore/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&MsgCreateOrder{}, "kiraHub/MsgCreateOrder", nil)
	cdc.RegisterConcrete(&MsgCreateOrderBook{}, "kiraHub/MsgCreateOrderBook", nil)
	cdc.RegisterConcrete(&MsgUpsertSignerKey{}, "kiraHub/MsgUpsertSignerKey", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgCreateOrder{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgCreateOrderBook{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpsertSignerKey{})
}

var (
	amino = codec.New()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}