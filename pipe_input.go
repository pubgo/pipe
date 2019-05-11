package pipe

import "reflect"

func ArrayOf(ps interface{}) *_func {
	_d := reflect.ValueOf(ps)

	var _ps []reflect.Value
	for i := 0; i < _d.Len(); i++ {
		_ps = append(_ps, _d.Index(i))
	}
	return &_func{
		params: _ps,
	}
}

func DataOf(ps ...interface{}) *_func {
	var vs []reflect.Value
	for _, v := range ps {
		vs = append(vs, reflect.ValueOf(v))
	}
	return &_func{
		params: vs,
	}
}

func SortBy(data interface{}, swap interface{}) *_func {
	return ArrayOf(data).SortBy(swap)
}
func Map(data interface{}, fn interface{}) *_func {
	return ArrayOf(data).Map(fn)
}

func MapExp(data interface{}, code string) *_func {
	return ArrayOf(data).MapExp(code)
}

func Reduce(data interface{}, code string) *_func {
	return ArrayOf(data).Reduce(code)
}

func Any(data interface{}, fn func(v interface{}) bool) bool {
	return ArrayOf(data).Any(fn)
}

func Every(data interface{}, fn func(v interface{}) bool) bool {
	return ArrayOf(data).Every(fn)
}

func Each(data interface{}, fn interface{}) {
	ArrayOf(data).Each(fn)
}

func FilterNil(data interface{}) *_func {
	return ArrayOf(data).FilterNil()
}

func FilterExp(data interface{}, filter string) *_func {
	return ArrayOf(data).FilterExp(filter)
}

func Filter(data interface{}, fn interface{}) *_func {
	return ArrayOf(data).Filter(fn)
}

func Pipe(data interface{}, fn interface{}) *_func {
	return ArrayOf(data).Pipe(fn)
}
