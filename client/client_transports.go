package client

import (
	"context"
	"github.com/go-openapi/runtime"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"

	httptransport "github.com/go-openapi/runtime/client"
)

// NewUnixSocketTransport creates a new clientTransport configured at the specified Unix socketPath.
func NewUnixSocketTransport(socketPath string, logger *logrus.Entry, debug bool) runtime.ClientTransport {
	socketTransport := &http.Transport{
		DialContext: func(ctx context.Context, network, path string) (net.Conn, error) {
			addr, err := net.ResolveUnixAddr("unix", socketPath)
			if err != nil {
				return nil, err
			}

			return net.DialUnix("unix", nil, addr)
		},
	}

	transport := httptransport.New(DefaultHost, DefaultBasePath, DefaultSchemes)
	transport.Transport = socketTransport

	if debug {
		transport.SetDebug(debug)
	}

	if logger != nil {
		transport.SetLogger(logger)
	}

	return transport
}
