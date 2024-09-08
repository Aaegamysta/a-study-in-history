package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dghubble/oauth1"
	"go.uber.org/zap"

	"github.com/aaegamysta/a-study-in-history/spec/pkg/events"
)

type Client struct {
	logger     *zap.SugaredLogger
	httpClient *http.Client
}

func New(ctx context.Context, logger *zap.SugaredLogger, cfg Config) *Client {
	oauth1Config := oauth1.NewConfig(cfg.ConsumerKey, cfg.ConsumerSecret)
	token := oauth1.NewToken(cfg.Token, cfg.TokenSecret)

	oauth1ConfiguredClient := oauth1Config.Client(context.Background(), token)
	return &Client{
		logger:     logger,
		httpClient: oauth1ConfiguredClient,
	}
}

func (c *Client) Post(ctx context.Context, e events.Event) (err error) {
	req := c.newRequest(ctx, e)
	res, err := c.httpClient.Do(req)
	if err != nil {
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if res.Body == nil {
		return fmt.Errorf("error posting tweet: response body is nil")
	}
	if res.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err = json.Unmarshal(b, &errorResponse)
		if err != nil {
			return fmt.Errorf("could not unmarshall error response after posting tweet, maybe something wrong with validation? : %s", err)
		}
		// TODO: persist error
		return fmt.Errorf("error posting tweet: %s", errorResponse.Detail)
	}
	var tweetResponse SuccessfulResponse
	err = json.Unmarshal(b, &tweetResponse)
	if err != nil {
		return fmt.Errorf("could not unmarshall tweet after tweeting, maybe something wrong with validation? : %v", err)
	}
	// TODO: persist tweet response
	return nil
}

func (c *Client) newRequest(ctx context.Context, e events.Event) *http.Request {
	// TODO: format event
	formattedEvent := e.FormatForPosting()
	tweet := map[string]interface{}{
		"text": formattedEvent,
	}

	jsonData, err := json.Marshal(tweet)
	if err != nil {
		c.logger.Panicf("something wrong happened while marshalling formatted event to json : %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.x.com/2/tweets", strings.NewReader(string(jsonData)))
	if err != nil {
		c.logger.Panicf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}
