package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	for i, test := range mergeTests {
		actual, err := merge(test.def, test.overrides)
		if err != nil {
			t.Errorf("Test #%d errored out: %v", i, err)
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Test #%dL Expected\n'%v'\nSaw\n'%v'\n", i, test.expected, actual)
		}
	}
}

var mergeTests = []struct {
	def       map[string]interface{}
	overrides map[string]interface{}
	expected  map[string]interface{}
}{
	{
		def:       map[string]interface{}{"foo": "bar"},
		overrides: map[string]interface{}{},
		expected:  map[string]interface{}{"foo": "bar"},
	},
	{
		def:       map[string]interface{}{"foo": "bar"},
		overrides: map[string]interface{}{"foo": "baz"},
		expected:  map[string]interface{}{"foo": "baz"},
	},
	{
		def:       map[string]interface{}{"foo": "bar"},
		overrides: map[string]interface{}{"foo": []string{"baz", "qux"}},
		expected:  map[string]interface{}{"foo": []string{"baz", "qux"}},
	},
	{
		def: map[string]interface{}{
			"foo": "bar",
			"bar": map[string]string{"key": "val"},
		},
		overrides: map[string]interface{}{
			"foo": "baz",
			"bar": map[string]string{"val": "key"},
		},
		expected: map[string]interface{}{
			"foo": "baz",
			"bar": map[string]string{
				"key": "val",
				"val": "key",
			},
		},
	},
}
