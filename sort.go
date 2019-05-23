package pipe

import (
	"reflect"
	"sort"
)

type reflectValueSlice struct {
	data []reflect.Value
	swap reflect.Value
}

func (c reflectValueSlice) Len() int {
	return len(c.data)
}
func (c reflectValueSlice) Swap(i, j int) {
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

func (c reflectValueSlice) Sort() []reflect.Value {
	sort.Sort(c)
	return c.data
}

func (c reflectValueSlice) Less(i, j int) bool {
	return c.swap.Call([]reflect.Value{c.data[i], c.data[j]})[0].Bool()
}
