package pipe

import (
	"github.com/antonmedv/expr"
	"github.com/pubgo/assert"
)

var _AssertFn = assert.AssertFn
var _FnOf = assert.FnOf
var _ST = assert.T
var _SWrap = assert.ErrWrap
var _IsNil = assert.IsNil
var _If = assert.If
var _IsPtr = assert.IsPtr

var _Eval = expr.Eval
