package chartserver

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	commonhttp "github.com/cowk8s/harbor/src/common/http"
)

const (
	clientTimeout         = 30 * time.Second
	maxIdleConnections    = 10
	idleConnectionTimeout = 30 * time.Second
)

var (
	once sync.Once
	chartTransport http.RoundTripper
)

// ChartClient is a http client to get the connect from the external http server
type ChartClient struct {
	// HTTP client
	httpClient *http.Client

	// Auth info
	credential *Credential
}

// NewChartClient is constructor of ChartClient
// Credential can be nil
func NewChartClient(credential *Credential) *ChartClient {
	once.Do(func() {
		chartTransport = commonhttp.NewTransport(
			commonhttp.WithMaxIdleConns(maxIdleConnections),
			commonhttp.WithIdleConnectionTimeout(idleConnectionTimeout),
		)
		
	})

	client := &http.Client{
		Timeout: clientTimeout,
		Transport: chartTransport,
	}

	return &ChartClient{
		httpClient: client,
		credential: credential,
	}
}

func (cc *ChartClient) GetContent(addr string) ([]byte, error) {
	response, err := cc.sendRequest(addr, http.MethodGet, nil)
	if err != nil {
		err = errors.Wrap(err, "get content field")
		return nil, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = errors.Wrap(err, "Read response body error")
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		text, err := extractError(content)
		if err != nil {
			err = errors.Wrap(err, "Extract content error failed")
			return nil, err
		}
		return nil, &commonhttp.Error{
			Code: response.StatusCode,
			Message:  text,
		}
	}
	return content, nil
}


