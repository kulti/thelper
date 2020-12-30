// Code generated by generator. DO NOT EDIT.

package t

import (
	"context"
	"testing"
)

func nonTestHelper(t int) {}

func helperWithoutHelper(t *testing.T) {} // want "test helper function should start from t.Helper()"

func helperWithHelper(t *testing.T) {
	t.Helper()
}

func helperWithEmptyStringBeforeHelper(t *testing.T) {

	t.Helper()
}

func helperWithHelperAfterAssignment(t *testing.T) { // want "test helper function should start from t.Helper()"
	_ = 0
	t.Helper()
}

func helperWithHelperAfterOtherCall(t *testing.T) { // want "test helper function should start from t.Helper()"
	f()
	t.Helper()
}

func helperWithHelperAfterOtherSelectionCall(t *testing.T) { // want "test helper function should start from t.Helper()"
	t.Fail()
	t.Helper()
}

func helperParamNotFirst(s string, i int, t *testing.T) { 
	t.Helper()
}

func helperParamSecondWithoutContext(s string, t *testing.T, i int) { 
	t.Helper()
}

func helperParamSecondWithContext(ctx context.Context, t *testing.T) {
	t.Helper()
}

func helperWithIncorrectName(o *testing.T) { 
	o.Helper()
}

func helperWithAnonymousHelper(t *testing.T) {
	t.Helper()
	func(t *testing.T) {}(t) // want "test helper function should start from t.Helper()"
}

func f() {}
