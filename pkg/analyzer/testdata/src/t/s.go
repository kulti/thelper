package t

import "testing"

func TestSubtestWithAnonymous(t *testing.T) {
	t.Run("sub", func(t *testing.T) {})
}

func TestSubtest(t *testing.T) {
	t.Run("sub", check)
}

func check(t *testing.T) {
}
