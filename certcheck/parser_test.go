package certcheck_test

import (
	"testing"

	"github.com/arthureichelberger/certcheck/certcheck"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()

	t.Run("it should return an error if the link is empty", func(t *testing.T) {
		_, err := certcheck.Parse("")
		require.Error(t, err)
	})

	t.Run("it should validate the link without the scheme", func(t *testing.T) {
		url, err := certcheck.Parse("example.com")
		require.NoError(t, err)
		require.Equal(t, "https://example.com", url.String())
	})
}
