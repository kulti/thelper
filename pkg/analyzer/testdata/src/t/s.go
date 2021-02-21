package t

import "testing"

func TestSubtestWithAnonymous(t *testing.T) {
	t.Run("sub", func(t *testing.T) {})
}

func TestSubtest(t *testing.T) {
	h := subhelper{}
	t.Run("sub", check)
	t.Run("sub sel", h.check)
}

func TestSubtestWithHelperAsRun(t *testing.T) {
	h := subhelper{}
	t.Run("sub with helper", anotherCheck)
	t.Run("sub sel", h.anotherCheck)
}

func check(t *testing.T) {
	anotherCheck(t)
}

func anotherCheck(t *testing.T) { // want "test helper function should start from t.Helper()"
}

type subhelper struct{}

func (sh subhelper) check(t *testing.T) {
	sh.anotherCheck(t)
}

func (sh subhelper) anotherCheck(t *testing.T) {} // want "test helper function should start from t.Helper()"
