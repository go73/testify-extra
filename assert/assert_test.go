package assert

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	t.Run("Equal net.IPs return true", func(t *testing.T) {
		// Set up
		mockT := &testing.T{}
		// Same IP address, different representation (requiring Equal())
		testExpected := net.IPv4zero.To4() // 4-byte representation
		testActual := net.IPv4zero         // 16-byte representation

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, actual)
	})

	t.Run("Unequal net.IPs return false", func(t *testing.T) {
		// Set up
		mockT := &testing.T{}
		testExpected := net.IPv4zero
		testActual := net.IPv6zero

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, actual)
	})

	t.Run("Equal time.Times return true", func(t *testing.T) {
		// Set up
		mockT := &testing.T{}
		testExpected := time.Unix(42, 0)
		testActual := testExpected.In(time.FixedZone("UTC-1", -1*60*60))

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, actual)
	})

	t.Run("Unequal time.Times return false", func(t *testing.T) {
		// Set up
		mockT := &testing.T{}
		testExpected := time.Unix(42, 0)
		testActual := testExpected.Add(1)

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, actual)
	})

	t.Run("Equal decimal.Decimal return true", func(t *testing.T) {
		// Set up
		mockT := &testing.T{}
		testExpected := decimal.New(1, 1)
		testActual := decimal.New(10, 0)

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, actual)
	})

	t.Run("Unequal decimal.Decimal return false", func(t *testing.T) {
		// Set up
		mockT := &testing.T{}
		testExpected := decimal.New(1_000_000, -6)
		testActual := decimal.New(1_000_001, -6)

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, actual)
	})
}

// go test -v -fuzz=Fuzz -fuzztime 30s github.com/go73/testify-extra/assert -fuzz FuzzEqualWithEqualInput
func FuzzEqualWithEqualInput(f *testing.F) {
	f.Add(int64(1361))
	f.Fuzz(func(t *testing.T, millis int64) {
		// Set up
		mockT := &testing.T{}
		testExpected := time.UnixMilli(millis)
		testActual := testExpected.In(time.FixedZone("UTC-1", -1*60*60))

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, actual)
	})
}

// go test -v -fuzz=Fuzz -fuzztime 30s github.com/go73/testify-extra/assert -fuzz FuzzEqualWithUnequalInput
func FuzzEqualWithUnequalInput(f *testing.F) {
	f.Add(int64(1361))
	f.Fuzz(func(t *testing.T, millis int64) {
		// Set up
		mockT := &testing.T{}
		testExpected := time.UnixMilli(millis)
		testActual := time.UnixMilli(millis + 1)

		// Test
		actual := Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, actual)
	})
}
