package errors

import "errors"

var (
	CreateClaudeClientBaseUrlError = errors.New("unsupport baseurl")
	ClaudeClientCallFormatError    = errors.New("response format parse error")
)
