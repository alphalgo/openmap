package openmap

import "testing"

func TestSearch(t *testing.T) {
	id := 1
	value := "hello"
	om := &Omap{
		m:  make(map[interface{}]interface{}),
		id: id,
		v:  value,
	}
	om.m[id] = value
	err := om.Search(id)

	assertErr(t, err, ErrNotFound)
}

func assertErr(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
