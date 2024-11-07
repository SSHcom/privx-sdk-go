package filters

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	ASC  = "ASC"  // ASC defines ascending sort direction.
	DESC = "DESC" // DESC defines descending sort direction.
)

// Option function type setting url values.
type Option func(*url.Values)

// Offset sets the offset for url values.
func Offset(offset int) Option {
	return func(v *url.Values) { v.Set("offset", strconv.Itoa(offset)) }
}

// Limit sets the limit for url values.
func Limit(limit int) Option {
	return func(v *url.Values) { v.Set("limit", strconv.Itoa(limit)) }
}

// Paging sets both the offset and limit for url values.
func Paging(offset, limit int) Option {
	return func(v *url.Values) {
		v.Set("offset", strconv.Itoa(offset))
		v.Set("limit", strconv.Itoa(limit))
	}
}

// Sort sets the sort key and direction for url values.
func Sort(key, dir string) Option {
	return func(v *url.Values) {
		v.Set("sortkey", key)
		v.Set("sortdir", dir)
	}
}

// SortAsc sets the sort key with ascending order for url values.
func SortAsc(key string) Option {
	return func(v *url.Values) {
		v.Set("sortkey", key)
		v.Set("sortdir", ASC)
	}
}

// SortDesc sets the sort key with descending order for url values.
func SortDesc(key string) Option {
	return func(v *url.Values) {
		v.Set("sortkey", key)
		v.Set("sortdir", DESC)
	}
}

// Filter sets the filter for url values.
func Filter(filter string) Option {
	return func(v *url.Values) { v.Set("filter", filter) }
}

// FuzzyCount sets the fuzzy count for url values.
func FuzzyCount(fuzzycount bool) Option {
	return func(v *url.Values) { v.Set("fuzzycount", strconv.FormatBool(fuzzycount)) }
}

// SetCustomParams set custom key-value parameter pairs freely.
func SetCustomParams(key, value string) Option {
	return func(v *url.Values) { v.Set(key, value) }
}

// SetStructParams convert struct to URL values and return Option from those values.
func SetStructParams(p interface{}) Option {
	return func(v *url.Values) {
		values, err := Values(p)
		if err != nil {
			fmt.Printf("failed converting struct into URL values: %v", err)
			return
		}

		options := urlValuesToOptions(values)
		for _, opt := range options {
			opt(v)
		}
	}
}

// urlValuesToOptions convert URL values into []Option.
func urlValuesToOptions(values url.Values) []Option {
	opts := make([]Option, 0)
	for key, vals := range values {
		for _, val := range vals {
			opts = append(opts, func(k, v string) Option {
				return func(params *url.Values) {
					params.Set(k, v)
				}
			}(key, val))
		}
	}
	return opts
}
