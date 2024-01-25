package text

import "testing"

func TestFirstNonEmpty(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		expect string
	}{
		{name: "no empty", args: []string{"foo", "bar"}, expect: "foo"},
		{name: "first non-empty", args: []string{"", "foo", "bar"}, expect: "foo"},
		{name: "second non-empty", args: []string{"", "", "bar"}, expect: "bar"},
		{name: "all empty", args: []string{"", "", ""}, expect: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonEmpty(tt.args...); got != tt.expect {
				t.Errorf("FirstNonEmpty() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestWithDefault(t *testing.T) {
	tests := []struct {
		name   string
		val    string
		def    string
		expect string
	}{
		{name: "val", val: "foo", def: "bar", expect: "foo"},
		{name: "def", val: "", def: "bar", expect: "bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDefault(tt.val, tt.def); got != tt.expect {
				t.Errorf("WithDefault() = %v, want %v", got, tt.expect)
			}
		})
	}
}
