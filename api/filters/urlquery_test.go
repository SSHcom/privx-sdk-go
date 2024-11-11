package filters

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestValues(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    url.Values
		wantErr bool
	}{
		{
			name: "Test Struct with omitempty tags no values",
			args: args{
				struct {
					StringTest string `url:"string_test,omitempty"`
					TimeTest   string `url:"time_test,omitempty"`
					IntTest    int    `url:"int_test,omitempty"`
					BoolTest   bool   `url:"bool_test,omitempty"`
				}{},
			},
			want:    url.Values{},
			wantErr: false,
		},
		{
			name: "Test Struct with omitempty with values",
			args: args{
				struct {
					StringTest string `url:"string_test,omitempty"`
					TimeTest   string `url:"time_test,omitempty"`
					IntTest    int    `url:"int_test,omitempty"`
					BoolTest   bool   `url:"bool_test,omitempty"`
				}{
					StringTest: "test",
					TimeTest:   time.Now().Format(time.RFC3339),
					BoolTest:   true,
					IntTest:    42,
				},
			},
			want: url.Values{
				"bool_test":   {"true"},
				"string_test": {"test"},
				"int_test":    {"42"},
				"time_test":   {time.Now().Format(time.RFC3339)},
			},
		},
		{
			name: "Test Struct with slice and custom tag options",
			args: args{
				struct {
					StringList []string `url:"string_list,comma"`
				}{
					StringList: []string{"item1", "item2", "item3"},
				},
			},
			want: url.Values{
				"string_list": {"item1,item2,item3"},
			},
			wantErr: false,
		},
		{
			name: "Test Struct with empty slice and omitempty",
			args: args{
				struct {
					StringList []string `url:"string_list,omitempty"`
				}{},
			},
			want:    url.Values{},
			wantErr: false,
		},
		{
			name: "Test Struct with pointer fields",
			args: args{
				struct {
					IntPtr *int `url:"int_ptr,omitempty"`
				}{
					IntPtr: func() *int { v := 100; return &v }(),
				},
			},
			want: url.Values{
				"int_ptr": {"100"},
			},
			wantErr: false,
		},
		{
			name: "Test Struct with nil pointer and omitempty",
			args: args{
				struct {
					IntPtr *int `url:"int_ptr,omitempty"`
				}{},
			},
			want:    url.Values{},
			wantErr: false,
		},
		{
			name: "Invalid type (non-struct input)",
			args: args{
				"this_is_not_a_struct",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty Test Struct without omitempty",
			args: args{
				struct {
					StringTest string `url:"string_test"`
					TimeTest   string `url:"time_test"`
					IntTest    int    `url:"int_test"`
					BoolTest   bool   `url:"bool_test"`
				}{},
			},
			want: url.Values{
				"string_test": {""},
				"time_test":   {""},
				"int_test":    {"0"},
				"bool_test":   {"false"},
			},
		},
		{
			name: "Test Struct with embedded struct",
			args: args{
				struct {
					EmbeddedStruct struct {
						SubStringTest string `url:"sub_string_test"`
						SubIntTest    int    `url:"sub_int_test"`
					}
					StringTest string `url:"string_test"`
					TimeTest   string `url:"time_test"`
					IntTest    int    `url:"int_test"`
					BoolTest   bool   `url:"bool_test"`
				}{
					EmbeddedStruct: struct {
						SubStringTest string `url:"sub_string_test"`
						SubIntTest    int    `url:"sub_int_test"`
					}{
						SubStringTest: "embedded",
						SubIntTest:    101,
					},
					StringTest: "parent",
					TimeTest:   time.Now().Format(time.RFC3339),
					BoolTest:   true,
					IntTest:    42,
				},
			},
			want: url.Values{
				"sub_string_test": {"embedded"},
				"sub_int_test":    {"101"},
				"bool_test":       {"true"},
				"string_test":     {"parent"},
				"int_test":        {"42"},
				"time_test":       {time.Now().Format(time.RFC3339)},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Values(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Values() error = %v, \nwantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
