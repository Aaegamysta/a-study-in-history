package twitter

type SuccessfulResponse struct {
	Data                TweetData `json:"data"`
	Text                string    `json:"text"`
	EditHistoryTweetIDs []string  `json:"edit_history_tweet_ids"`
}

type TweetData struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}
