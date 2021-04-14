package exchange

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAmount(t *testing.T) {
	e := Exch{}

	amount, errExch := e.ConvertedAmount(1, "RON", "2021-04-13")
	require.Nil(t, errExch, "passed currency is not correct")
	require.Equal(t, 4.9223, amount)

	amount, errExch = e.ConvertedAmount(1, "RON", "2021-04-12")
	require.Nil(t, errExch, "passed currency is not correct")
	require.Equal(t, 4.9203, amount)
}
