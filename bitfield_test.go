package bitfield

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBitfield(t *testing.T) {
	b := New(16) // 128 bit field

	require.Nil(t, b.Set(0))
	require.Nil(t, b.Set(1))
	require.Nil(t, b.Set(100))
	require.Nil(t, b.Set(102))
	require.NotNil(t, b.Set(128))
	require.NotNil(t, b.Set(-1))

	b = New(16)

	for i := 0; i < 128; i++ {
		if i%3 == 0 {
			continue
		}
		require.Nil(t, b.Set(i))
	}

	for i := 0; i < 128; i++ {
		if i%3 == 0 {
			assert.False(t, b.Has(i))
		} else {
			assert.True(t, b.Has(i))
		}
	}
}
