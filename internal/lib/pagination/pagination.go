// Package pagination contains everything related to pagination and is a support subdomain
package pagination

import "strconv"

const (
	maxItemsByPage = 10
)

func New(pageStr, limitStr string) *Pagination {
	limit, err := strconv.ParseUint(limitStr, 10, 8)
	if err != nil {
		limit = maxItemsByPage
	}

	if limit > maxItemsByPage {
		limit = maxItemsByPage
	}

	page, _ := strconv.ParseUint(pageStr, 10, 32)

	return &Pagination{page: uint32(page), limit: uint8(limit)}
}

// Pagination manage the records pagination using a number page and limit for each query
type Pagination struct {
	page         uint32
	limit        uint8
	totalResults uint64
}

// SetTotalResults sets the total results related to the query performed
func (p *Pagination) SetTotalResults(totalResults uint64) {
	p.totalResults = totalResults
}

// TotalResults returns the total results to the query performed
func (p *Pagination) TotalResults() uint64 {
	return p.totalResults
}

// SetPage sets the page
func (p *Pagination) SetPage(page uint32) {
	p.page = page
}

// Page returns the current page
func (p *Pagination) Page() uint32 {
	return p.page
}

func (p *Pagination) Pages() uint64 {
	if p.totalResults == 0 {
		return 0
	}

	return p.totalResults / uint64(p.limit)
}

// SetLimit sets the limit records for query
func (p *Pagination) SetLimit(limit uint8) {
	p.limit = limit
}

// Limit returns the limit records
func (p *Pagination) Limit() uint8 {
	return p.limit
}

// Start returns the initial index
func (p *Pagination) Start() uint64 {
	return uint64(p.page) * uint64(p.limit)
}

// Stop returns the last index
func (p *Pagination) Stop() uint64 {
	return uint64(p.page)*uint64(p.limit) + uint64(p.limit)
}
