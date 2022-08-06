package require

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
	"time"
)

type MockT struct {
	Failed bool
}

func (t *MockT) FailNow() {
	t.Failed = true
}

func (t *MockT) Errorf(format string, args ...interface{}) {
	_, _ = format, args
}

func TestEqual(t *testing.T) {
	t.Run("Equal net.IPs return true", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		// Same IP address, different representation (requiring Equal())
		testExpected := net.IPv4zero.To4() // 4-byte representation
		testActual := net.IPv4zero         // 16-byte representation

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})

	t.Run("Unequal net.IPs return false", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := net.IPv4zero
		testActual := net.IPv6zero

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})

	t.Run("Equal time.Times return true", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := time.Unix(42, 0)
		testActual := testExpected.In(time.FixedZone("UTC-1", -1*60*60))

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})

	t.Run("Unequal time.Times return false", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := time.Unix(42, 0)
		testActual := testExpected.Add(1)

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})

	t.Run("Equal decimal.Decimal return true", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := decimal.New(1, 1)
		testActual := decimal.New(10, 0)

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})

	t.Run("Unequal decimal.Decimal return false", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := decimal.New(1_000_000, -6)
		testActual := decimal.New(1_000_001, -6)

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})
}

// go test -v -fuzz=Fuzz -fuzztime 30s github.com/go73/testify-extra/require -fuzz FuzzEqualWithEqualInput
func FuzzEqualWithEqualInput(f *testing.F) {
	f.Add(int64(1361))
	f.Fuzz(func(t *testing.T, millis int64) {
		// Set up
		mockT := new(MockT)
		testExpected := time.UnixMilli(millis)
		testActual := testExpected.In(time.FixedZone("UTC-1", -1*60*60))

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})
}

// go test -v -fuzz=Fuzz -fuzztime 30s github.com/go73/testify-extra/require -fuzz FuzzEqualWithUnequalInput
func FuzzEqualWithUnequalInput(f *testing.F) {
	f.Add(int64(1361))
	f.Fuzz(func(t *testing.T, millis int64) {
		// Set up
		mockT := new(MockT)
		testExpected := time.UnixMilli(millis)
		testActual := time.UnixMilli(millis + 1)

		// Test
		Equal(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})
}

func TestNotEqual(t *testing.T) {
	t.Run("Equal net.IPs return true", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		// Same IP address, different representation (requiring Equal())
		testExpected := net.IPv4zero.To4() // 4-byte representation
		testActual := net.IPv4zero         // 16-byte representation

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})

	t.Run("Unequal net.IPs return false", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := net.IPv4zero
		testActual := net.IPv6zero

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})

	t.Run("Equal time.Times return true", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := time.Unix(42, 0)
		testActual := testExpected.In(time.FixedZone("UTC-1", -1*60*60))

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})

	t.Run("Unequal time.Times return false", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := time.Unix(42, 0)
		testActual := testExpected.Add(1)

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})

	t.Run("Equal decimal.Decimal return true", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := decimal.New(1, 1)
		testActual := decimal.New(10, 0)

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})

	t.Run("Unequal decimal.Decimal return false", func(t *testing.T) {
		// Set up
		mockT := new(MockT)
		testExpected := decimal.New(1_000_000, -6)
		testActual := decimal.New(1_000_001, -6)

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})
}

// go test -v -fuzz=Fuzz -fuzztime 30s github.com/go73/testify-extra/require -fuzz FuzzEqualWithEqualInput
func FuzzNotEqualWithEqualInput(f *testing.F) {
	f.Add(int64(1361))
	f.Fuzz(func(t *testing.T, millis int64) {
		// Set up
		mockT := new(MockT)
		testExpected := time.UnixMilli(millis)
		testActual := testExpected.In(time.FixedZone("UTC-1", -1*60*60))

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.True(t, mockT.Failed)
	})
}

// go test -v -fuzz=Fuzz -fuzztime 30s github.com/go73/testify-extra/require -fuzz FuzzEqualWithUnequalInput
func FuzzNotEqualWithUnequalInput(f *testing.F) {
	f.Add(int64(1361))
	f.Fuzz(func(t *testing.T, millis int64) {
		// Set up
		mockT := new(MockT)
		testExpected := time.UnixMilli(millis)
		testActual := time.UnixMilli(millis + 1)

		// Test
		NotEqual(mockT, testExpected, testActual)

		// Verify
		require.False(t, mockT.Failed)
	})
}
