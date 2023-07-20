package pb

func (p *Request) Compatible() {
	if len(p.Body) == 0 && len(p.Entity) > 0 {
		p.Body = p.Entity
	}

	if p.URL == "" && p.Uri != "" {
		p.URL = p.Uri
	}
}
