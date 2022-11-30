package bitarray

import (
	amino "github.com/tendermint/tendermint2/pkgs/amino"
	bitarraypb "github.com/tendermint/tendermint2/pkgs/bitarray/pb"
	proto "google.golang.org/protobuf/proto"
)

func (goo BitArray) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *bitarraypb.BitArray
	{
		if IsBitArrayReprEmpty(goo) {
			var pbov *bitarraypb.BitArray
			msg = pbov
			return
		}
		pbo = new(bitarraypb.BitArray)
		{
			pbo.Bits = int64(goo.Bits)
		}
		{
			goorl := len(goo.Elems)
			if goorl == 0 {
				pbo.Elems = nil
			} else {
				pbos := make([]uint64, goorl)
				for i := 0; i < goorl; i += 1 {
					{
						goore := goo.Elems[i]
						{
							pbos[i] = uint64(goore)
						}
					}
				}
				pbo.Elems = pbos
			}
		}
	}
	msg = pbo
	return
}

func (goo BitArray) EmptyPBMessage(cdc *amino.Codec) (msg proto.Message) {
	pbo := new(bitarraypb.BitArray)
	msg = pbo
	return
}

func (goo *BitArray) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *bitarraypb.BitArray = msg.(*bitarraypb.BitArray)
	{
		if pbo != nil {
			{
				(*goo).Bits = int(int(pbo.Bits))
			}
			{
				var pbol int = 0
				if pbo.Elems != nil {
					pbol = len(pbo.Elems)
				}
				if pbol == 0 {
					(*goo).Elems = nil
				} else {
					goors := make([]uint64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Elems[i]
							{
								pboev := pboe
								goors[i] = uint64(pboev)
							}
						}
					}
					(*goo).Elems = goors
				}
			}
		}
	}
	return
}

func (_ BitArray) GetTypeURL() (typeURL string) {
	return "/tm.BitArray"
}

func IsBitArrayReprEmpty(goor BitArray) (empty bool) {
	{
		empty = true
		{
			if goor.Bits != int(0) {
				return false
			}
		}
		{
			if len(goor.Elems) != 0 {
				return false
			}
		}
	}
	return
}
