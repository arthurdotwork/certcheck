package main

import (
	"context"
	"fmt"
	"os"

	"github.com/arthureichelberger/certcheck/certcheck"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx); err != nil {
		log.Error().Err(err).Msg("failed to run certcheck")
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	if len(os.Args) != 2 {
		return fmt.Errorf("missing argument")
	}

	addr, err := certcheck.Parse(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	certs, err := certcheck.Check(ctx, addr)
	if err != nil {
		return fmt.Errorf("failed to check certificates: %w", err)
	}

	for _, cert := range certs {
		fmt.Printf("%s expires on %s, in %s.\n", cert.CommonName, cert.NotAfter.Format("2006-01-02"), cert.FormatExpirationDate())
	}

	return nil
}
