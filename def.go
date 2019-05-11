package pipe

type IPipe interface {
	SortBy(swap interface{}) *_func

	Pipe(fn interface{}) *_func

	P(tags ...string)

	MapExp(code string) *_func
	Map(fn interface{}) *_func

	Reduce(fn interface{}) *_func

	Any(fn func(v interface{}) bool) bool
	Every(fn func(v interface{}) bool) bool
	Each(fn interface{})

	FilterNil() *_func
	FilterExp(filter string) *_func
	Filter(fn func(v interface{}) bool) *_func

	MustNotNil()
}
