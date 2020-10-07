package adtype

import (
	"reflect"
	"sort"
)

func BoolStreamOf(arg ...bool) BoolStream {
	return arg
}
func BoolStreamFrom(arg []bool) BoolStream {
	return arg
}

func (self *BoolStream) Add(arg bool) *BoolStream {
	return self.AddAll(arg)

}

func (self *BoolStream) AddAll(arg ...bool) *BoolStream {
	*self = append(*self, arg...)
	return self
}

func (self *BoolStream) AddSafe(arg *bool) *BoolStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *BoolStream) AllMatch(fn func(arg bool, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *BoolStream) AnyMatch(fn func(arg bool, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *BoolStream) Clone() *BoolStream {
	temp := make([]bool, self.Len())
	copy(temp, *self)
	return (*BoolStream)(&temp)
}

func (self *BoolStream) Copy() *BoolStream {
	return self.Clone()
}

func (self *BoolStream) Concat(arg []bool) *BoolStream {
	return self.AddAll(arg...)
}

func (self *BoolStream) Delete(index int) *BoolStream {
	return self.DeleteRange(index, index)
}

func (self *BoolStream) DeleteRange(startIndex int, endIndex int) *BoolStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *BoolStream) Equals(arr []bool) bool {
	if (*self == nil) != (arr == nil) || len(*self) != len(arr) {
		return false
	}
	for i := range *self {
		if (*self)[i] != arr[i] {
			return false
		}
	}
	return true
}
func (self *BoolStream) Filter(fn func(arg bool, index int) bool) *BoolStream {
	_array := BoolStreamOf()
	self.ForEach(func(v bool, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *BoolStream) Find(fn func(arg bool, index int) bool) *bool {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *BoolStream) FindIndex(fn func(arg bool, index int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *BoolStream) First() *bool {
	return self.Get(0)
}

func (self *BoolStream) ForEach(fn func(arg bool, index int)) *BoolStream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *BoolStream) ForEachRight(fn func(arg bool, index int)) *BoolStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *BoolStream) GroupBy(fn func(arg bool, index int) string) map[string][]bool {
	m := map[string][]bool{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *BoolStream) GroupByValues(fn func(arg bool, index int) string) [][]bool {
	tmp := [][]bool{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *BoolStream) IndexOf(arg bool) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *BoolStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *BoolStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *BoolStream) Last() *bool {
	return self.Get(self.Len() - 1)
}

func (self *BoolStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *BoolStream) Limit(limit int) *BoolStream {
	self.Slice(0, limit)
	return self
}
func (self *BoolStream) Map(fn func(arg bool, index int) bool) *BoolStream {
	return self.ForEach(func(v bool, i int) { self.Set(i, fn(v, i)) })
}

func (self *BoolStream) MapAny(fn func(arg bool, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) Map2Int(fn func(arg bool, index int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) Map2Int32(fn func(arg bool, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) Map2Int64(fn func(arg bool, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) Map2Float32(fn func(arg bool, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) Map2Float64(fn func(arg bool, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) Map2Bytes(fn func(arg bool, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) Map2String(fn func(arg bool, index int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *BoolStream) NoneMatch(fn func(arg bool, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *BoolStream) Get(index int) *bool {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *BoolStream) Peek(fn func(arg *bool, index int)) *BoolStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *BoolStream) Reduce(fn func(result, current bool, index int) bool) *BoolStream {
	return self.ReduceInit(fn, false)
}
func (self *BoolStream) ReduceInit(fn func(result, current bool, index int) bool, initialValue bool) *BoolStream {
	result := BoolStreamOf()
	self.ForEach(func(v bool, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *BoolStream) Reverse() *BoolStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *BoolStream) Replace(fn func(arg bool, index int) bool) *BoolStream {
	return self.Map(fn)
}

func (self *BoolStream) Set(index int, val bool) *BoolStream {
	if len(*self) > index {
		(*self)[index] = val
	}
	return self
}

func (self *BoolStream) Skip(skip int) *BoolStream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *BoolStream) Slice(startIndex int, n int) *BoolStream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []bool{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *BoolStream) Sort(fn func(i, j int) bool) *BoolStream {
	sort.Slice(*self, fn)
	return self
}

func (self *BoolStream) SortStable(fn func(i, j int) bool) *BoolStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *BoolStream) ToList() []bool {
	return self.Val()
}

func (self *BoolStream) Val() []bool {
	if self == nil {
		return []bool{}
	}
	return *self
}

func (self *BoolStream) While(fn func(arg bool, index int) bool) *BoolStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}
func (self *BoolStream) Contains(arg bool) bool {
	return self.FindIndex(func(_arg bool, index int) bool { return _arg == arg }) != -1
}
func (self *BoolStream) Distinct() *BoolStream {
	m := map[bool]bool{}
	self.Filter(func(arg bool, index int) bool {
		_, ok := m[arg]
		if !ok {
			m[arg] = true
		}
		return !ok
	})
	return self
}
