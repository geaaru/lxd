// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// See https://github.com/golang/go/blob/master/CONTRIBUTORS
// Licensed under the same terms as Go itself:
// https://github.com/golang/go/blob/master/LICENSE

//go:build go1.7
// +build go1.7

package subtest

import "testing"

// Run runs function f as a subtest of t.
func Run(t *testing.T, name string, f func(t *testing.T)) {
	t.Run(name, f)
}
