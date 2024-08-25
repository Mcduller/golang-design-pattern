package cd_driver

import (
	"mediator_szl"
	"mediator_szl/domain"
	"mediator_szl/types"
)

type normalCdDriver struct {
	Data types.Data
}

func (n *normalCdDriver) ReadDataV2(data types.Data) {
	n.Data = data
	mediator_szl.GetMediatorInstance().Change(n)
}

func (n *normalCdDriver) GetData() types.Data {
	return n.Data
}

func NewCDDriver() domain.CDDriver {
	return &normalCdDriver{}
}
