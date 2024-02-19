package certcheck

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func Check(ctx context.Context, addr *url.URL) ([]Certificate, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, addr.String(), nil)
	if err != nil {
		return []Certificate{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return []Certificate{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	certs := make([]Certificate, 0, len(resp.TLS.PeerCertificates))
	for _, cert := range resp.TLS.PeerCertificates {
		certs = append(certs, Certificate{
			CommonName: cert.Subject.CommonName,
			NotAfter:   cert.NotAfter,
		})
	}

	return certs, nil
}
