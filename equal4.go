// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package equal4 provides functionality for comparison of an error messages.
package equal4

import (
	"regexp"
	"strings"
)

func ErrorEqual(err, err2 error) bool {
	if err != nil && err2 == nil {
		return false
	}

	if err == nil && err2 != nil {
		return false
	}

	if err != nil && err2 != nil && err.Error() != err2.Error() {
		return false
	}

	return true
}

func ErrorContains(err error, str string) bool {
	if err != nil && str == "" {
		return false
	}

	if err == nil && str != "" {
		return false
	}

	if err != nil && str != "" && !strings.Contains(err.Error(), str) {
		return false
	}

	return true
}

func ErrorMatch(err error, re *regexp.Regexp) bool {
	if err != nil && re == nil {
		return false
	}

	if err == nil && re != nil {
		return false
	}

	if err != nil && re != nil && !re.MatchString(err.Error()) {
		return false
	}

	return true
}
