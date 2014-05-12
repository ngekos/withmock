// Copyright 2013 Julian Phillips.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"os"
)

type Err struct {
	Ctxt string
	Err  error
}

func (c Err) Error() string {
	return c.Err.Error()
}

func (c Err) Context() string {
	if c2, ok := c.Err.(Err); ok {
		return c.Ctxt + ":" + c2.Context()
	} else {
		return c.Ctxt
	}
}

func IsNotExist(err error) bool {
	if e, ok := err.(Err); ok {
		return IsNotExist(e.Err)
	}
	return os.IsNotExist(err)
}