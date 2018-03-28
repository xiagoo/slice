package slice

import (
	"sort"
	"reflect"
	"fmt"
)

type functions struct {
	length int
	less   func(i, j int) bool
	swap   func(i, j int)
}

func (f *functions) Len() int           { return f.length }
func (f *functions) Less(i, j int) bool { return f.less(i, j) }
func (f *functions) Swap(i, j int)      { f.swap(i, j) }

func Sort(slice interface{}, less func(i, j int) bool) {
	sort.Sort(SortSlice(slice, less))
}

func SortSlice(slice interface{}, less func(i, j int) bool) sort.Interface {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("slice.Sort need slice value of type %T", slice))
	}
	return &functions{
		length: sv.Len(),
		less:   less,
		swap:   Swapper(sv),
	}
}

func Swapper(v reflect.Value) func(i, j int) {
	tmp := reflect.New(v.Type().Elem()).Elem()
	return func(i, j int) {
		v1 := v.Index(i)
		v2 := v.Index(j)
		tmp.Set(v1)
		v1.Set(v2)
		v2.Set(tmp)
	}
}