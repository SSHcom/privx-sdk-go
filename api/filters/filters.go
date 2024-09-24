package filters

const (
	ASC  = "ASC"  // ASC defines ascending sort direction.
	DESC = "DESC" // DESC defines descending sort direction.
)

// Params defines optional query parameters.
type Params struct {
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Sortkey string `json:"sortkey,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
}

// Default constructor for default filters.
func Default() Params {
	return Params{
		Offset:  0,
		Limit:   50,
		Sortdir: "ASC",
	}
}

// Option function type modifying Params.
type Option func(p *Params)

// Offset sets the offset for Params.
func Offset(o int) Option {
	return func(p *Params) { p.Offset = o }
}

// Limit sets the limit for Params.
func Limit(l int) Option {
	return func(p *Params) { p.Limit = l }
}

// Paging sets both the offset and limit for Params.
func Paging(o, l int) Option {
	return func(p *Params) {
		p.Offset = o
		p.Limit = l
	}
}

// Sort sets the sort key and direction for Params.
func Sort(key, dir string) Option {
	return func(p *Params) {
		p.Sortkey = key
		p.Sortdir = dir
	}
}

// SortAsc sets the sort key with ascending order for Params.
func SortAsc(key string) Option {
	return func(p *Params) {
		p.Sortkey = key
		p.Sortdir = ASC
	}
}

// SortDesc sets the sort key with descending order for Params.
func SortDesc(key string) Option {
	return func(p *Params) {
		p.Sortkey = key
		p.Sortdir = DESC
	}
}
