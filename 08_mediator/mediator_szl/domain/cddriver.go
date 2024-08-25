package domain

import "mediator_szl/types"

type CDDriver interface {
	ReadDataV2(data types.Data)
	GetData() types.Data
}
