package common

import (
	"github.com/shopspring/decimal"
	"net"
	"time"
)

// These just needs to compile to give us assurance that our target types
// conform to the Equatable[T] interface
var _ Equatable[net.IP] = (net.IP)(nil)
var _ Equatable[time.Time] = (*time.Time)(nil)
var _ Equatable[decimal.Decimal] = (*decimal.Decimal)(nil)
