package proxy

import (
	"strconv"
	"strings"
	"time"

	"github.com/xuender/kit/base"
	"github.com/xuender/kit/counter"
	"github.com/xuender/kit/logs"
	"github.com/xuender/limit"
)

type Limits struct {
	counts *counter.Counter[string]
	limits map[string]*limit.Mutex
	qps    map[string]int
}

func NewLimits() *Limits {
	ret := &Limits{
		limits: map[string]*limit.Mutex{},
		counts: counter.NewCounter[string](),
		qps:    map[string]int{},
	}

	go ret.reset()

	return ret
}

func (p *Limits) reset() {
	timer := time.NewTicker(time.Minute)

	for range timer.C {
		p.log()
		p.counts.Clean()
	}
}

func (p *Limits) log() {
	buf := strings.Builder{}

	for key, qps := range p.qps {
		if num := p.counts.Get(key); num > 0 {
			if buf.Len() > 0 {
				buf.WriteString(", ")
			}

			buf.WriteString(key)
			buf.WriteRune(':')
			buf.WriteString(strconv.Itoa(int(num)))
			buf.WriteString("[QPS]")
			buf.WriteString(strconv.Itoa(int(num / base.Sixty)))
			buf.WriteString("/")
			buf.WriteString(strconv.Itoa(qps))
		}
	}

	if buf.Len() > 0 {
		logs.I.Printf("limits: %s\n", buf.String())
	}
}

func (p *Limits) Add(key string, qps int) *Limits {
	timeOut := time.Second * base.Ten

	p.qps[key] = qps
	p.limits[key] = limit.NewMutex(qps, timeOut)

	return p
}

func (p *Limits) Check(url string) error {
	for key := range p.qps {
		if Has(url, key) {
			if err := p.limits[key].Wait(); err != nil {
				return err
			}

			p.counts.Inc(key)

			return nil
		}
	}

	return nil
}

func Has(url, key string) bool {
	if url == "" {
		return false
	}

	index := strings.Index(url, key)

	if index < 0 {
		return false
	}

	if key[0] == '/' {
		if last := index + len(key); len(url) > last {
			return url[last] == '?' || url[last] == '/'
		}
	}

	return true
}
