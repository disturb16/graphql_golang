package settings

type contextCey string

const (
	// RequestTracking is a map containing
	// all variables related to kong gateway
	// like Kong-Request-ID, api-key
	RequestTracking contextCey = "request-tracking"
)
