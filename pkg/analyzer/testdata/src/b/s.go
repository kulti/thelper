package b

import "testing"

func BenchmarkSubtestWithAnonymous(b *testing.B) {
	b.Run("sub", func(b *testing.B) {})
}

func BenchmarkSubtest(b *testing.B) {
	b.Run("sub", check)
}

func check(b *testing.B) {
}
