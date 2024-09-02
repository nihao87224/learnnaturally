package commands

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"net/http"
	"net/url"
	"os"
)

func MyGeminiClient(ctx context.Context, proxyURL string) (*genai.Client, error) {
	key := os.Getenv("API_KEY")

	var clientOpts []option.ClientOption
	if len(proxyURL) > 0 {
		c := &http.Client{Transport: &proxyRoundTripper{
			APIKey:   key,
			ProxyURL: proxyURL,
		}}

		clientOpts = append(clientOpts, option.WithHTTPClient(c))
	} else {
		clientOpts = append(clientOpts, option.WithAPIKey(key))
	}

	client, err := genai.NewClient(ctx, clientOpts...)
	return client, err
}

type proxyRoundTripper struct {
	// APIKey is the API Key to set on requests.
	APIKey string

	// ProxyURL is the URL of the proxy server. If empty, no proxy is used.
	ProxyURL string
}

func (t *proxyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()

	if t.ProxyURL != "" {
		proxyURL, err := url.Parse(t.ProxyURL)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	newReq := req.Clone(req.Context())
	vals := newReq.URL.Query()
	vals.Set("key", t.APIKey)
	newReq.URL.RawQuery = vals.Encode()

	resp, err := transport.RoundTrip(newReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
