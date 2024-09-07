package mstdn

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
	"go.uber.org/zap"
)

type Client struct {
	httpClient *http.Client
	logger     *zap.SugaredLogger
	cfg        Config
}

func New(ctx context.Context, logger *zap.SugaredLogger, cfg Config) *Client {
	return &Client{
		httpClient: &http.Client{},
		logger:     logger,
		cfg:        cfg,
	}
}

// TODO: define type response
func (c *Client) Post(ctx context.Context, e events.Event) (err error) {
	req := c.newRequest(ctx, e)
	res, err := c.httpClient.Do(req)
	if err != nil {
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if res.Body == nil {
		return fmt.Errorf("error posting status: response body is nil")
	}
	if res.StatusCode != http.StatusOK {
		var postFailureResponse PostFailureResponse 
		err = json.Unmarshal(b, &postFailureResponse)
		if err != nil {
			return fmt.Errorf("could not unmarshall error response after posting status, maybe something wrong with validation? : %s", err)
		}
		// TODO: persist error
		return fmt.Errorf("error posting status: %s", postFailureResponse.Error)
	}
	var postSuccessResponse PostSuccessResponse 
	err = json.Unmarshal(b, &postSuccessResponse)
	if err != nil {
		return fmt.Errorf("could not unmarshall status after posting, maybe something wrong with validation? : %v", err)
	}
	// TODO: persist tweet response
	return nil
}

func (c *Client) newRequest(ctx context.Context, e events.Event) *http.Request {
	formattedEvent := e.FormatForPosting()
	status := StatusPost{
		Status: formattedEvent,
		// TODO: add media ids via upload, maybe?
		MediaIds: []string{},
	}

	jsonData, err := json.Marshal(status)
	if err != nil {
		c.logger.Panicf("something wrong happened while marshalling formatted event to json : %v", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", c.cfg.Endpoint, strings.NewReader(string(jsonData)))
	if err != nil {
		c.logger.Panicf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.cfg.AccessToken))
	return req
}
