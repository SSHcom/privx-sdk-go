package filters

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	encoderType = reflect.TypeOf(new(Encoder)).Elem()
	timeType    = reflect.TypeOf(time.Time{})
)

// Encoder is an interface implemented by any type that wishes to encode
// itself into URL values in a non-standard way.
type Encoder interface {
	EncodeValues(key string, v *url.Values) error
}

// tagOptions is the string following a comma in a struct field's "url" tag, or
// the empty string. It does not include the leading comma.
type tagOptions []string

// Contains checks whether the tagOptions contains the specified option.
func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}

// Values convert struct into URL values by reflecting over its fields and tags.
func Values(s interface{}) (url.Values, error) {
	v := reflect.ValueOf(s)
	values := make(url.Values)

	// Dereference pointers, return empty map if nil.
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return values, nil
		}
		v = v.Elem()
	}

	if s == nil {
		return values, nil
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got: %T", s)
	}

	err := reflectValue(values, v, "")

	return values, err
}

// reflectValue recursively populates URL values by reflecting over
// fields of a struct, handling embedded structs, slices, arrays, and custom types.
func reflectValue(values url.Values, val reflect.Value, scope string) error {
	embedded := []reflect.Value{}
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		// skip unexported fields
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}

		sv := val.Field(i)
		tag := sf.Tag.Get("url")
		// skip ignored fields
		if tag == "-" {
			continue
		}

		// split url tag into its base name and options
		name, opts := parseTag(tag)
		if name == "" {
			// check for embedded struct fields
			if sf.Anonymous {
				// ensure value is the actual struct and not a pointer to it
				v := reflect.Indirect(sv)
				if v.IsValid() && v.Kind() == reflect.Struct {
					// save embedded struct for later processing
					embedded = append(embedded, v)
					continue
				}
			}

			name = sf.Name
		}

		// skip if omitempty is set and passed value is empty
		if opts.Contains("omitempty") && isEmptyValue(sv) {
			continue
		}

		// check if value implements Encoder interface for custom encoding
		if sv.Type().Implements(encoderType) {
			// if sv is a nil pointer and the custom encoder is defined on a non-pointer
			// method receiver, set sv to the zero value of the underlying type
			if !reflect.Indirect(sv).IsValid() && sv.Type().Elem().Implements(encoderType) {
				sv = reflect.New(sv.Type().Elem())
			}

			m := sv.Interface().(Encoder)
			if err := m.EncodeValues(name, &values); err != nil {
				return err
			}
			continue
		}

		// recursively dereference pointers. break on nil pointers
		for sv.Kind() == reflect.Ptr {
			// stop if nil to prevent further dereferencing
			if sv.IsNil() {
				break
			}
			sv = sv.Elem()
		}

		switch sv.Kind() {
		case reflect.Slice, reflect.Array:
			if sv.Len() == 0 {
				// skip if slice or array is empty
				continue
			}
			handleSliceOrArray(values, name, sv, opts, sf)
			continue
		case reflect.Struct:
			// handle recursively nested structs
			if err := reflectValue(values, sv, name); err != nil {
				return err
			}
			continue
		}
		values.Add(name, valueString(sv, opts, sf))
	}

	// Handle embedded structs
	for _, f := range embedded {
		if err := reflectValue(values, f, scope); err != nil {
			return err
		}
	}

	return nil
}

// handleSliceOrArray processes slices/arrays in a struct field, either joining them with a delimiter
// or adding each element separately to the URL values map.
func handleSliceOrArray(values url.Values, name string, sv reflect.Value, opts tagOptions, sf reflect.StructField) {
	var delimiter string
	switch {
	case opts.Contains("comma"):
		delimiter = ","
	case opts.Contains("space"):
		delimiter = " "
	case opts.Contains("semicolon"):
		delimiter = ";"
	case opts.Contains("brackets"):
		name += "[]"
	default:
		delimiter = sf.Tag.Get("del")
	}

	if delimiter != "" {
		var b strings.Builder
		for i := 0; i < sv.Len(); i++ {
			if i > 0 {
				b.WriteString(delimiter)
			}
			b.WriteString(valueString(sv.Index(i), opts, sf))
		}
		values.Add(name, b.String())
	} else {
		for i := 0; i < sv.Len(); i++ {
			k := name
			if opts.Contains("numbered") {
				k = fmt.Sprintf("%s%d", name, i)
			}
			values.Add(k, valueString(sv.Index(i), opts, sf))
		}
	}
}

// valueString returns the string representation of a value based on its type and tag options.
func valueString(v reflect.Value, opts tagOptions, sf reflect.StructField) string {
	// get underlying value, return empty string for invalid values
	v = reflect.Indirect(v)
	if !v.IsValid() {
		return ""
	}

	if v.Kind() == reflect.Bool && opts.Contains("int") {
		return strconv.Itoa(convertBool(v.Bool()))
	}

	if v.Type() == timeType {
		return formatTimeValue(v.Interface().(time.Time), opts, sf)
	}

	return fmt.Sprint(v.Interface())
}

// parseTag split struct field's "url" tag into its name and optional comma separated options.
func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

// convertBool convert a boolean value to 1 (true) or 0 (false).
func convertBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

// formatTimeValue formats a time.Time value based on provided tag options or struct field tag.
// Defaults to RFC3339 format if no specific option is provided.
func formatTimeValue(t time.Time, opts tagOptions, sf reflect.StructField) string {
	if opts.Contains("unix") {
		return strconv.FormatInt(t.Unix(), 10)
	}
	if opts.Contains("unixmilli") {
		return strconv.FormatInt(t.UnixNano()/1e6, 10)
	}
	if opts.Contains("unixnano") {
		return strconv.FormatInt(t.UnixNano(), 10)
	}
	if layout := sf.Tag.Get("layout"); layout != "" {
		return t.Format(layout)
	}
	return t.Format(time.RFC3339)
}

// isEmptyValue checks if a value should be considered empty for the purposes
// of omitting fields with the "omitempty" option.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	if z, ok := v.Interface().(interface{ IsZero() bool }); ok {
		return z.IsZero()
	}

	return false
}
