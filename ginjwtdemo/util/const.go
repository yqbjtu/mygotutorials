package util

import "errors"

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token is not active yet")
	TokenMalformed   error = errors.New("Malformed token")
	TokenInvalid     error = errors.New("can't handle this token")
)

const (
	// 这个是需要保密的一段信息
	SignKey         string = "a87x80wfebei90f8532f16f423b125616dea9b75"
	Gin_Context_Key string = "claims"
)
