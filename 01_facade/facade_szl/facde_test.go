package facade_szl

import (
	"testing"
)

func Test_facadeImpl_Run(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{
			name: "1",
			want: "ARun,BRun",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFacade()
			if got := f.Run(); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
