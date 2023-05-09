package csv

import "sync"

type results struct {
	values map[string]int
	sync.Mutex
}

func NewResult() *results {
	return &results{
		values: make(map[string]int),
	}
}

// Add adds a domain count to the results
func (r *results) Add(domain string) {
	r.Lock()
	r.values[domain]++
	r.Unlock()
}
