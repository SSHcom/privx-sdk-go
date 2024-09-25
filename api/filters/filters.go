package filters

const (
	ASC  = "ASC"  // ASC defines ascending sort direction.
	DESC = "DESC" // DESC defines descending sort direction.
)

// Params defines optional query parameters.
type Params struct {
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	SortKey string `json:"sortkey,omitempty"`
	SortDir string `json:"sortdir,omitempty"`
}

// Default constructor for default filters.
func Default() Params {
	return Params{
		Offset:  0,
		Limit:   50,
		SortDir: "ASC",
	}
}

// Option function type modifying Params.
type Option func(p *Params)

// Offset sets the offset for Params.
func Offset(offset int) Option {
	return func(p *Params) { p.Offset = offset }
}

// Limit sets the limit for Params.
func Limit(limit int) Option {
	return func(p *Params) { p.Limit = limit }
}

// Paging sets both the offset and limit for Params.
func Paging(offset, limit int) Option {
	return func(p *Params) {
		p.Offset = offset
		p.Limit = limit
	}
}

// Sort sets the sort key and direction for Params.
func Sort(key, dir string) Option {
	return func(p *Params) {
		p.SortKey = key
		p.SortDir = dir
	}
}

// SortAsc sets the sort key with ascending order for Params.
func SortAsc(key string) Option {
	return func(p *Params) {
		p.SortKey = key
		p.SortDir = ASC
	}
}

// SortDesc sets the sort key with descending order for Params.
func SortDesc(key string) Option {
	return func(p *Params) {
		p.SortKey = key
		p.SortDir = DESC
	}
}
