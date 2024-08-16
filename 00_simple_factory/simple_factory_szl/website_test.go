package simple_factory_szl

import (
	"reflect"
	"simple_factory_szl/const"
	"testing"
)

func TestNewWebSite(t *testing.T) {
	type args struct {
		siteType _const.WebSiteType
	}
	tests := []struct {
		name string
		args args
		want _const.WebAddress
	}{
		{
			name: "baidu",
			args: args{siteType: _const.WebSiteTypeBaidu},
			want: _const.BaiduWebSiteAddress,
		},
		{
			name: "google",
			args: args{siteType: _const.WebSiteTypeGoogle},
			want: _const.GoogleWebSiteAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWebSite(tt.args.siteType).GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
