// Copyright © 2017 ben dewan <benj.dewan@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
