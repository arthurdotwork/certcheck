package certcheck_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/arthureichelberger/certcheck/certcheck"
	"github.com/stretchr/testify/require"
)

func TestChecker(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t.Run("it should return an error if it can not do the request", func(t *testing.T) {
		_, err := certcheck.Check(ctx, &url.URL{Scheme: "http", Host: "localhost:9999"})
		require.Error(t, err)
	})

	t.Run("it should return the certificates", func(t *testing.T) {
		certs, err := certcheck.Check(ctx, &url.URL{Scheme: "https", Host: "example.com"})
		require.NoError(t, err)
		require.NotEmpty(t, certs)
	})
}
