package code

type StatusCode int

const (
	OK                 StatusCode = 200200
	NOT_AUTH           StatusCode = 200401
	NOT_FOUND          StatusCode = 200404
	METHOD_NOT_ALLOWED StatusCode = 200405
	SERVER_ERROR       StatusCode = 200500
	RELEASE_ERROR      StatusCode = 200501
)
