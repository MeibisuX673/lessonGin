package model

type Query struct {
	Page         uint
	Filters      []string
	Orders       []string
	RangeFilters []string
	Limit        uint
}
