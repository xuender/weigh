package pb

import "net/http"

func (p *Response) Compatible(rid int32) {
	p.Id = rid
	p.Code = p.StatusCode
	p.Entity = p.Body
	p.Body = nil
}

func NewErr(err error) *Response {
	return &Response{
		Error:      err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}
