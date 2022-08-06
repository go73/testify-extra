Testify Extra
=============

Extras to the stretchr/testify test library.

Delegating Equal()/NotEqual()
-----------------------------
Testify's `Equal()` implementation uses
[`reflect.DeepEqual(...)`](https://github.com/stretchr/testify/blob/v1.8.0/assert/assertions.go#L66)
to check for equality. This makes sense when comparing most values, however
there are certain types that are semantically equivalent even though they are
not strictly equal using this method.

In such cases, the de facto standard is to use an `Equal(T) bool` method to
define semantic equivalence. This is used within the Go standard library:

- [net.IP](https://pkg.go.dev/net#IP.Equal)
- [time.Time](https://pkg.go.dev/time#Time.Equal)

And is also adopted by some third-party libraries:

- [decimal.Decimal](https://pkg.go.dev/github.com/shopspring/decimal#Decimal.Equal)

`assert` example shown here, but the `require` version is available as well:
```go
package yours

import (
	"testing"
	"time"
	"github.com/go37/testify-extra/assert"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t,
		time.UnixMilli(123),
		time.UnixMilli(123).In(time.FixedZone("UTC-1", -1*60*60)),
		"they should be equal")

	// assert inequality
	assert.NotEqual(t,
		time.UnixMilli(123), time.UnixMilli(456),
		"they should not be equal")

}
```

Installation
============

    go get github.com/go73/testify-extra
