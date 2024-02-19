package certcheck_test

import (
	"testing"
	"time"

	"github.com/arthureichelberger/certcheck"
	"github.com/stretchr/testify/require"
)

func TestFormatExpirationDate(t *testing.T) {
	t.Parallel()

	t.Run("it should return the expiration date in days", func(t *testing.T) {
		c := certcheck.Certificate{NotAfter: time.Now().Add(48 * time.Hour)}
		require.Equal(t, "2 days", c.FormatExpirationDate())
	})

	t.Run("it should return the expiration date in hours", func(t *testing.T) {
		c := certcheck.Certificate{NotAfter: time.Now().Add(2 * time.Hour)}
		require.Equal(t, "2 hours", c.FormatExpirationDate())
	})

	t.Run("it should return the expiration date in minutes", func(t *testing.T) {
		c := certcheck.Certificate{NotAfter: time.Now().Add(2 * time.Minute)}
		require.Equal(t, "2 minutes", c.FormatExpirationDate())
	})

	t.Run("it should return the expiration date in a minute", func(t *testing.T) {
		c := certcheck.Certificate{NotAfter: time.Now().Add(1 * time.Minute)}
		require.Equal(t, "1 minute", c.FormatExpirationDate())
	})
}
