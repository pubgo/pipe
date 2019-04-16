package pipe

import (
	_assert "github.com/pubgo/assert"
)

var assert = _assert.Bool

func If(b bool, tv, fv interface{}) interface{} {
	if b {
		return tv
	}
	return fv
}
