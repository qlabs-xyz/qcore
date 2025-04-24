package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	Amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// cdc.RegisterConcrete(&MsgDistributeAuctionReward{}, "gemmint/MsgDistributeAuctionReward", nil)
	// cdc.RegisterConcrete(&MsgMintRedeemedOptionNFT{}, "gemmint/MsgMintRedeemedOptionNFT", nil)
	// // this line is used by starport scaffolding # 2

	// // proposals
	// cdc.RegisterConcrete(&UpdateAnnualProvisionsProposal{}, "gemchain/update-annual-provisions-proposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {

	// registry.RegisterImplementations((*sdk.Msg)(nil),
	// 	&MsgDistributeAuctionReward{},
	// 	&MsgMintRedeemedOptionNFT{},
	// )
	// this line is used by starport scaffolding # 3

	// msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	// proposals
	// registry.RegisterImplementations(
	// 	(*govtypes.Content)(nil),
	// 	&UpdateAnnualProvisionsProposal{},
	// )

}
