package restark_test

import (
	"testing"

	"github.com/influx6/restark"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	var stack restark.StackApp
	require.Error(t, stack.Apply(1))
}
