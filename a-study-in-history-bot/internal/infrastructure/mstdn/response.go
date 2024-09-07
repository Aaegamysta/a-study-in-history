package mstdn

type PostSuccessResponse struct {
	ID string `json:"id"`
	CreatedAt string `json:"created_at"`

}

type PostFailureResponse struct {
	Error string `json:"error"`
}