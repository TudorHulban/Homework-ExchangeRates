package exchange

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeed(t *testing.T) {
	e := Exch{}

	_, errExch := e.ConvertedAmount(1.00, "RON", "2021-04-13")
	require.Nil(t, errExch, "passed currency is not correct")
}
