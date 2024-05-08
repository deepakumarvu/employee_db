package apierror

type APIError struct {
	HttpStatusCode int    `json:"-"`
	ErrCode        int    `json:"code,omitempty"`
	ErrorMessage   string `json:"error_msg,omitempty"`
}

func (e *APIError) Error() string {
	return e.ErrorMessage
}

var _ error = (*APIError)(nil)
