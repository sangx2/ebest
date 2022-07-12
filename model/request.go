package model

type Request struct {
	ResName string

	InBlocks []interface{}

	IsOccurs bool

	RespChan chan *Response `json:"-"`
}

func NewQueryRequest(resName string, isOccurs bool, inBlock ...interface{}) *Request {
	return &Request{
		ResName:  resName,
		InBlocks: inBlock,
		IsOccurs: isOccurs,
		RespChan: make(chan *Response, 1),
	}
}
