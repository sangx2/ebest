package model

type Response struct {
	OutBlocks []interface{}

	Error error
}

func NewResponse(outBlocks []interface{}, err error) *Response {
	return &Response{
		OutBlocks: outBlocks,
		Error:     err,
	}
}
