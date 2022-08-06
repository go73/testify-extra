package assert

import (
	"fmt"
	"github.com/go73/testify-extra/internal/common"
	"github.com/stretchr/testify/assert"
)

func Equal[T common.Equatable[T]](t assert.TestingT, expected, actual T, msgAndArgs ...interface{}) bool {
	if !expected.Equal(actual) {
		return assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %v\n"+
			"actual  : %v", expected, actual), msgAndArgs...)
	}

	return true
}

func NotEqual[T common.Equatable[T]](t assert.TestingT, expected, actual T, msgAndArgs ...interface{}) bool {
	if expected.Equal(actual) {
		return assert.Fail(t, fmt.Sprintf("Should not be: %#v\n", actual), msgAndArgs...)
	}

	return true
}
