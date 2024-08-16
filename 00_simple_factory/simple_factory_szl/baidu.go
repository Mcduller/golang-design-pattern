package simple_factory_szl

import (
	"simple_factory_szl/const"
)

type baiduImpl struct {
}

func NewBaiduImpl() WebSite {
	return baiduImpl{}
}

func (b baiduImpl) GetAddress() _const.WebAddress {
	return _const.BaiduWebSiteAddress
}
