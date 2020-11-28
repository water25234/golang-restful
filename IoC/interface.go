package ioc

// API means
type API interface {
	// Call means
	Call(httpMethod, path string, requestBody interface{}) (response interface{}, err error)
}

// Handle means
type Handle interface {
	// Handle means
	Handle() (response interface{}, err error)
}
