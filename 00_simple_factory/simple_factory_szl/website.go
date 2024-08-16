package simple_factory_szl

import (
	"simple_factory_szl/const"
)

type WebSite interface {
	GetAddress() _const.WebAddress
}

func NewWebSite(siteType _const.WebSiteType) WebSite {
	switch siteType {
	case _const.WebSiteTypeGoogle:
		return NewGoogle()
	case _const.WebSiteTypeBaidu:
		return NewBaiduImpl()
	}
	return nil
}
