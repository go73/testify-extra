package require

import (
	"fmt"
	"github.com/go73/testify-extra/internal/common"
	"github.com/stretchr/testify/require"
)

func Equal[T common.Equatable[T]](t require.TestingT, expected, actual T, msgAndArgs ...interface{}) {
	if !expected.Equal(actual) {
		require.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %v\n"+
			"actual  : %v", expected, actual), msgAndArgs...)
	}
}
