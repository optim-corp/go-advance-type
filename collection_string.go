package adtype

import (
	"reflect"
	"sort"
)

func StringStreamOf(arg ...string) StringStream {
	return arg
}
func StringStreamFrom(arg []string) StringStream {
	return arg
}

func (self *StringStream) Add(arg string) *StringStream {
	return self.AddAll(arg)

}

func (self *StringStream) AddAll(arg ...string) *StringStream {
	*self = append(*self, arg...)
	return self
}

func (self *StringStream) AddSafe(arg *string) *StringStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *StringStream) AllMatch(fn func(arg string, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *StringStream) AnyMatch(fn func(arg string, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *StringStream) Clone() *StringStream {
	temp := make([]string, self.Len())
	copy(temp, *self)
	return (*StringStream)(&temp)
}

func (self *StringStream) Copy() *StringStream {
	return self.Clone()
}

func (self *StringStream) Concat(arg []string) *StringStream {
	return self.AddAll(arg...)
}

func (self *StringStream) Delete(index int) *StringStream {
	return self.DeleteRange(index, index)
}

func (self *StringStream) DeleteRange(startIndex int, endIndex int) *StringStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *StringStream) Equals(arr []string) bool {
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
func (self *StringStream) Filter(fn func(arg string, index int) bool) *StringStream {
	_array := StringStreamOf()
	self.ForEach(func(v string, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *StringStream) Find(fn func(arg string, index int) bool) *string {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *StringStream) FindIndex(fn func(arg string, index int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *StringStream) First() *string {
	return self.Get(0)
}

func (self *StringStream) ForEach(fn func(arg string, index int)) *StringStream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *StringStream) ForEachRight(fn func(arg string, index int)) *StringStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *StringStream) GroupBy(fn func(arg string, index int) string) map[string][]string {
	m := map[string][]string{}

	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *StringStream) GroupByValues(fn func(arg string, index int) string) [][]string {
	tmp := [][]string{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *StringStream) IndexOf(arg string) int {
	for index, _arg := range *self {
		if reflect.DeepEqual(_arg, arg) {
			return index
		}
	}
	return -1
}
func (self *StringStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *StringStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *StringStream) Last() *string {
	return self.Get(self.Len() - 1)
}

func (self *StringStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *StringStream) Limit(limit int) *StringStream {
	self.Slice(0, limit)
	return self
}
func (self *StringStream) Map(fn func(arg string, index int) string) *StringStream {
	return self.ForEach(func(v string, i int) { self.Set(i, fn(v, i)) })
}

func (self *StringStream) NoneMatch(fn func(arg string, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *StringStream) Get(index int) *string {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *StringStream) Peek(fn func(arg *string, index int)) *StringStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *StringStream) Reduce(fn func(result, current string, index int) string) *StringStream {
	return self.ReduceInit(fn, "")
}
func (self *StringStream) ReduceInit(fn func(result, current string, index int) string, initialValue string) *StringStream {
	result := StringStreamOf()
	self.ForEach(func(v string, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *StringStream) Reverse() *StringStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *StringStream) Replace(fn func(arg string, index int) string) *StringStream {
	return self.Map(fn)
}

func (self *StringStream) Set(index int, val string) *StringStream {
	if len(*self) > index {
		(*self)[index] = val
	}
	return self
}

func (self *StringStream) Skip(skip int) *StringStream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *StringStream) Slice(startIndex int, n int) *StringStream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []string{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *StringStream) Sort(fn func(i, j int) bool) *StringStream {
	sort.Slice(*self, fn)
	return self
}

func (self *StringStream) SortStable(fn func(i, j int) bool) *StringStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *StringStream) ToList() []string {
	return self.Val()
}

func (self *StringStream) Val() []string {
	if self == nil {
		return []string{}
	}
	return *self
}

func (self *StringStream) While(fn func(arg string, index int) bool) *StringStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *StringStream) MapAny(fn func(arg string, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Int(fn func(arg string, index int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Int32(fn func(arg string, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Int64(fn func(arg string, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Float32(fn func(arg string, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Float64(fn func(arg string, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Bool(fn func(arg string, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Bytes(fn func(arg string, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Max(fn func(arg string, index int) string) string {
	return MaxStr(self.Val()...)
}
func (self *StringStream) Min(fn func(arg string, index int) string) string {
	return MinStr(self.Val()...)
}

func (self *StringStream) Sum() string {
	self.Reduce(func(result, current string, index int) string {
		return result + current
	})
	if v := self.Last(); v != nil {
		return *v
	}
	return ""
}
func (self *StringStream) Contains(arg string) bool {
	return self.FindIndex(func(_arg string, index int) bool { return _arg == arg }) != -1
}
func (self *StringStream) Distinct() *StringStream {
	m := map[string]bool{}
	self.Filter(func(arg string, index int) bool {
		_, ok := m[arg]
		if !ok {
			m[arg] = true
		}
		return !ok
	})
	return self
}
