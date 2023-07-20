package pb

import "github.com/xuender/weigh/utils"

func (p *Config) IsSerial(url string) bool {
	if len(p.Serial) == 0 {
		return false
	}

	for _, key := range p.Serial {
		if utils.Has(url, key) {
			return true
		}
	}

	return false
}

func (p *Config) Group(reqs []*Request) ([]*Request, []*Request) {
	async := []*Request{}
	serial := []*Request{}

	for _, req := range reqs {
		if p.IsSerial(req.URL) {
			serial = append(serial, req)
		} else {
			async = append(async, req)
		}
	}

	return async, serial
}
