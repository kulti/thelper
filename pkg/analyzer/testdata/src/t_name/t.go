// Code generated by generator. DO NOT EDIT.

package t

import "testing"

func nonTestHelper(t int) {}

func helperWithoutHelper(t *testing.T) {} 

func helperWithHelper(t *testing.T) {
	t.Helper()
}

func helperWithEmptyStringBeforeHelper(t *testing.T) {

	t.Helper()
}

func helperWithHelperAfterAssignment(t *testing.T) { 
	_ = 0
	t.Helper()
}

func helperWithHelperAfterOtherCall(t *testing.T) { 
	f()
	t.Helper()
}

func helperWithHelperAfterOtherSelectionCall(t *testing.T) { 
	t.Fail()
	t.Helper()
}

func helperWithNotFirst(s string, t *testing.T, i int) { 
	t.Helper()
}

func helperWithIncorrectName(o *testing.T) { // want "parameter \\*testing.T should have name t"
	o.Helper()
}

func f() {}
