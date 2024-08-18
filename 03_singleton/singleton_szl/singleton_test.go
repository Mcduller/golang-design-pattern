package singleton_szl

import "testing"

func TestSingleton_Foo(t *testing.T) {
	t.Run("get one", func(t *testing.T) {
		want := "Singleton Szl"
		newSingleton := NewSingleton()
		if got := newSingleton.Foo(); got != want {
			t.Errorf("want:%s,got:%s", want, got)
		}
	})
}
