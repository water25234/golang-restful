package api

// HTTPClient means
type HTTPClient interface {
	Get(path string) (resp interface{}, err error)

	Post(path string, requestBody interface{}) (resp interface{}, err error)

	Put(path string, requestBody interface{}) (resp interface{}, err error)

	Patch(path string, requestBody interface{}) (resp interface{}, err error)

	Delete(path string, requestBody interface{}) (resp interface{}, err error)
}
