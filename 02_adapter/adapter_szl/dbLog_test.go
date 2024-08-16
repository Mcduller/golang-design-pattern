package adapter_szl

import "testing"

func TestGetLogFromDB(t *testing.T) {
	want := "From File"
	dbLog := NewDBLog()
	if got := dbLog.GetLogFromDB(); got != want {
		t.Errorf("want:%s,got:%s", want, got)
	}
}
