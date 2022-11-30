package hd

import (
	amino "github.com/tendermint/tendermint2/pkgs/amino"
	hdpb "github.com/tendermint/tendermint2/pkgs/crypto/hd/pb"
	proto "google.golang.org/protobuf/proto"
)

func (goo BIP44Params) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *hdpb.Bip44Params
	{
		if IsBip44ParamsReprEmpty(goo) {
			var pbov *hdpb.Bip44Params
			msg = pbov
			return
		}
		pbo = new(hdpb.Bip44Params)
		{
			pbo.Purpose = uint32(goo.Purpose)
		}
		{
			pbo.CoinType = uint32(goo.CoinType)
		}
		{
			pbo.Account = uint32(goo.Account)
		}
		{
			pbo.Change = bool(goo.Change)
		}
		{
			pbo.AddressIndex = uint32(goo.AddressIndex)
		}
	}
	msg = pbo
	return
}

func (goo BIP44Params) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(hdpb.Bip44Params)
	msg = pbo
	return
}

func (goo *BIP44Params) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *hdpb.Bip44Params = msg.(*hdpb.Bip44Params)
	{
		if pbo != nil {
			{
				(*goo).Purpose = uint32(pbo.Purpose)
			}
			{
				(*goo).CoinType = uint32(pbo.CoinType)
			}
			{
				(*goo).Account = uint32(pbo.Account)
			}
			{
				(*goo).Change = bool(pbo.Change)
			}
			{
				(*goo).AddressIndex = uint32(pbo.AddressIndex)
			}
		}
	}
	return
}

func (_ BIP44Params) GetTypeURL() (typeURL string) {
	return "/tm.Bip44Params"
}

func IsBip44ParamsReprEmpty(goor BIP44Params) (empty bool) {
	{
		empty = true
		{
			if goor.Purpose != uint32(0) {
				return false
			}
		}
		{
			if goor.CoinType != uint32(0) {
				return false
			}
		}
		{
			if goor.Account != uint32(0) {
				return false
			}
		}
		{
			if goor.Change != bool(false) {
				return false
			}
		}
		{
			if goor.AddressIndex != uint32(0) {
				return false
			}
		}
	}
	return
}
