// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package equal4_test

import (
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"testing"

	"github.com/danil/equal4"
)

func line() int { _, _, l, _ := runtime.Caller(1); return l }

var ErrorEqualTestCases = []struct {
	line      int
	errors    []error
	expected  bool
	benchmark bool
}{
	{
		line:     line(),
		errors:   []error{nil, nil},
		expected: true,
	},
	{
		line:     line(),
		errors:   []error{errors.New("foo"), errors.New("bar")},
		expected: false,
	},
	{
		line:     line(),
		errors:   []error{errors.New("xyz"), errors.New("xyz")},
		expected: true,
	},
	{
		line:     line(),
		errors:   []error{errors.New("something went wrong"), nil},
		expected: false,
	},
	{
		line:     line(),
		errors:   []error{nil, errors.New("something went wrong")},
		expected: false,
	},
}

func TestErrorEqual(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range ErrorEqualTestCases {
		tc := tc
		t.Run(fmt.Sprint(tc.errors), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			ok := equal4.EqualErrors(tc.errors[0], tc.errors[1])
			if ok != tc.expected {
				t.Errorf("unexpected error equality, expected: %t, recieved: %t %s", tc.expected, ok, linkToExample)
			}
		})
	}
}

var ErrorContainsTestCases = []struct {
	line      int
	error     error
	string    string
	expected  bool
	benchmark bool
}{
	{
		line:     line(),
		error:    nil,
		string:   "",
		expected: true,
	},
	{
		line:     line(),
		error:    errors.New("foo"),
		string:   "bar",
		expected: false,
	},
	{
		line:     line(),
		error:    errors.New("foobar"),
		string:   "bar",
		expected: true,
	},
	{
		line:     line(),
		error:    errors.New("xyz"),
		string:   "xyz",
		expected: true,
	},
	{
		line:     line(),
		error:    errors.New("something went wrong"),
		string:   "",
		expected: false,
	},
	{
		line:     line(),
		error:    nil,
		string:   "something went wrong",
		expected: false,
	},
}

func TestErrorContains(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range ErrorContainsTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%s %s", tc.error, tc.string), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			ok := equal4.ErrorContains(tc.error, tc.string)
			if ok != tc.expected {
				t.Errorf("unexpected error contains, expected: %t, recieved: %t %s", tc.expected, ok, linkToExample)
			}
		})
	}
}

var ErrorMatchTestCases = []struct {
	line      int
	error     error
	regexp    *regexp.Regexp
	expected  bool
	benchmark bool
}{
	{
		line:     line(),
		error:    nil,
		regexp:   nil,
		expected: true,
	},
	{
		line:     line(),
		error:    errors.New("foo"),
		regexp:   regexp.MustCompile("^bar$"),
		expected: false,
	},
	{
		line:     line(),
		error:    errors.New("foobar"),
		regexp:   regexp.MustCompile("bar"),
		expected: true,
	},
	{
		line:     line(),
		error:    errors.New("xyz"),
		regexp:   regexp.MustCompile("xyz"),
		expected: true,
	},
	{
		line:     line(),
		error:    errors.New("foo"),
		regexp:   nil,
		expected: false,
	},
	{
		line:     line(),
		error:    nil,
		regexp:   regexp.MustCompile("foo"),
		expected: false,
	},
}

func TestErrorMatch(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)
	for _, tc := range ErrorMatchTestCases {
		tc := tc
		t.Run(fmt.Sprintf("%s %s", tc.error, tc.regexp), func(t *testing.T) {
			t.Parallel()
			linkToExample := fmt.Sprintf("%s:%d", testFile, tc.line)

			ok := equal4.ErrorMatch(tc.error, tc.regexp)
			if ok != tc.expected {
				t.Errorf("unexpected error match, expected: %t, recieved: %t %s", tc.expected, ok, linkToExample)
			}
		})
	}
}
