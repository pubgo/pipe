package pipe

import (
	"encoding/json"
	"reflect"
)

func (t *_func) ToRaw() []reflect.Value {
	return t.params
}

func (t *_func) ToString() string {
	return t.ToJson()
}

func (t *_func) ToData(fn ...interface{}) interface{} {
	var _t reflect.Type
	for _, _v := range t.params {
		if _v.IsValid() {
			_t = _v.Type()
			break
		}
	}

	if _t == nil {
		return nil
	}

	for i := 0; i < len(t.params); i++ {
		if !t.params[i].IsValid() {
			t.params[i] = reflect.Zero(_t)
		}
	}

	_rst := reflect.MakeSlice(reflect.SliceOf(_t), 0, len(t.params))
	_rst = reflect.Append(_rst, t.params...)

	if len(fn) != 0 && !_IsNil(fn[0]) && reflect.TypeOf(fn[0]).Kind() == reflect.Func {
		reflect.ValueOf(fn[0]).Call([]reflect.Value{_rst})
		return nil
	}

	return _rst.Interface()
}
func IsNil(p interface{}) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = false
		}
	}()

	if !reflect.ValueOf(p).IsValid() {
		return true
	}

	return reflect.ValueOf(p).IsNil()
}

func (t *_func) ToJson() string {
	var _res []interface{}
	for _, _p := range t.params {
		_res = append(_res, _If(IsNil(_p),"", _FnOf(_p.Interface)))
	}


	dt, err := json.Marshal(_res)
	_SWrap(err, "data json error")

	return string(dt)
}
