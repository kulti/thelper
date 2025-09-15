package t

import (
	"testing"
	"testing/synctest"
)

func TestSynctestWithAnonymous(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {})
}

func TestSynctest(t *testing.T) {
	h := subhelper{}
	synctest.Test(t, check)
	synctest.Test(t, h.check)
}

func TestSynctestWithHelperAsTest(t *testing.T) {
	h := subhelper{}
	synctest.Test(t, anotherCheck)
	synctest.Test(t, h.anotherCheck)
}

func TestSynctestWithBuilder(t *testing.T) {
	synctest.Test(t, subtestBuilder("first"))
	synctest.Test(t, subtestBuilder("second"))
	synctest.Test(t, subtestBuilderAnotherFile())
	synctest.Test(t, func() func(t *testing.T) { return func(t *testing.T) {} }())

	fixture := subtestFixture{}
	synctest.Test(t, fixture.subtest())
}
