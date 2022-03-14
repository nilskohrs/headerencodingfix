// Package headerencodingfix is a plugin
package headerencodingfix

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
)

// Config the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// headerEncodingFix a Traefik plugin.
type headerEncodingFix struct {
	next    http.Handler
	encoder *encoding.Encoder
}

// New creates a new headerEncodingFix plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	encoding, err := ianaindex.MIME.Encoding("latin1")
	if err != nil {
		log.Fatal(err)
	}

	return &headerEncodingFix{
		next:    next,
		encoder: encoding.NewEncoder(),
	}, nil
}

func (c *headerEncodingFix) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, values := range req.Header {
		for i := 0; i < len(values); i++ {
			val, err := c.encoder.String(values[i])
			if err != nil {
				log.Fatal(err)
			}
			values[i] = val
		}
	}
}
