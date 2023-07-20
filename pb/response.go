package pb

func (p *Response) Compatible(rid int32) {
	p.Id = rid
	p.Code = p.StatusCode
	p.Entity = p.Body
	p.Body = nil
}
