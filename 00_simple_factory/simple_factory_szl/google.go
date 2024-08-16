package simple_factory_szl

import (
	"simple_factory_szl/const"
)

type googleImpl struct {
}

func NewGoogle() WebSite {
	return googleImpl{}
}

func (g googleImpl) GetAddress() _const.WebAddress {
	return _const.GoogleWebSiteAddress
}
