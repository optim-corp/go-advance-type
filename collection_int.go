package adtype

import (
	"sort"
)

func IntegerStreamOf(arg ...int) IntegerStream {
	return arg
}
func IntegerStreamFrom(arg []int) IntegerStream {
	return arg
}

func (self *IntegerStream) Add(arg int) *IntegerStream {
	return self.AddAll(arg)

}

func (self *IntegerStream) AddAll(arg ...int) *IntegerStream {
	*self = append(*self, arg...)
	return self
}

func (self *IntegerStream) AddSafe(arg *int) *IntegerStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *IntegerStream) AllMatch(fn func(arg int, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *IntegerStream) AnyMatch(fn func(arg int, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *IntegerStream) Clone() *IntegerStream {
	temp := make([]int, self.Len())
	copy(temp, *self)
	return (*IntegerStream)(&temp)
}

func (self *IntegerStream) Copy() *IntegerStream {
	return self.Clone()
}

func (self *IntegerStream) Concat(arg []int) *IntegerStream {
	return self.AddAll(arg...)
}

func (self *IntegerStream) Delete(index int) *IntegerStream {
	return self.DeleteRange(index, index)
}

func (self *IntegerStream) DeleteRange(startIndex int, endIndex int) *IntegerStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *IntegerStream) Equals(arr []int) bool {
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
func (self *IntegerStream) Filter(fn func(arg int, index int) bool) *IntegerStream {
	_array := IntegerStreamOf()
	self.ForEach(func(v int, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *IntegerStream) Find(fn func(arg int, index int) bool) *int {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *IntegerStream) FindIndex(fn func(arg int, index int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *IntegerStream) First() *int {
	return self.Get(0)
}

func (self *IntegerStream) ForEach(fn func(arg int, index int)) *IntegerStream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *IntegerStream) ForEachRight(fn func(arg int, index int)) *IntegerStream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *IntegerStream) GroupBy(fn func(arg int, index int) string) map[string][]int {
	m := map[string][]int{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *IntegerStream) GroupByValues(fn func(arg int, index int) string) [][]int {
	tmp := [][]int{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *IntegerStream) IndexOf(arg int) int {
	for index, _arg := range *self {
		if _arg == arg {
			return index
		}
	}
	return -1
}
func (self *IntegerStream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *IntegerStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *IntegerStream) Last() *int {
	return self.Get(self.Len() - 1)
}

func (self *IntegerStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *IntegerStream) Limit(limit int) *IntegerStream {
	self.Slice(0, limit)
	return self
}
func (self *IntegerStream) Map(fn func(arg int, index int) int) *IntegerStream {
	return self.ForEach(func(v int, i int) { self.Set(i, fn(v, i)) })
}

func (self *IntegerStream) MapAny(fn func(arg int, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2Float32(fn func(arg int, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2Int32(fn func(arg int, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2Int64(fn func(arg int, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2Float64(fn func(arg int, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2Bool(fn func(arg int, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2Bytes(fn func(arg int, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2String(fn func(arg int, index int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) NoneMatch(fn func(arg int, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *IntegerStream) Get(index int) *int {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *IntegerStream) Peek(fn func(arg *int, index int)) *IntegerStream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *IntegerStream) Reduce(fn func(result, current int, index int) int) *IntegerStream {
	return self.ReduceInit(fn, 0)
}
func (self *IntegerStream) ReduceInit(fn func(result, current int, index int) int, initialValue int) *IntegerStream {
	result := IntegerStreamOf()
	self.ForEach(func(v int, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *IntegerStream) Reverse() *IntegerStream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *IntegerStream) Replace(fn func(arg int, index int) int) *IntegerStream {
	return self.Map(fn)
}

func (self *IntegerStream) Set(index int, val int) *IntegerStream {
	if len(*self) > index {
		(*self)[index] = val
	}
	return self
}

func (self *IntegerStream) Skip(skip int) *IntegerStream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *IntegerStream) Slice(startIndex int, n int) *IntegerStream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []int{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *IntegerStream) Sort(fn func(i, j int) bool) *IntegerStream {
	sort.Slice(*self, fn)
	return self
}

func (self *IntegerStream) SortStable(fn func(i, j int) bool) *IntegerStream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *IntegerStream) ToList() []int {
	return self.Val()
}

func (self *IntegerStream) Val() []int {
	if self == nil {
		return []int{}
	}
	return *self
}

func (self *IntegerStream) While(fn func(arg int, index int) bool) *IntegerStream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *IntegerStream) Min() int {
	return MinInt(self.Val()...)
}

func (self *IntegerStream) Max() int {
	return MaxInt(self.Val()...)
}

func (self *IntegerStream) Sum() int {
	return SumInt(self.Val()...)
}

func (self *IntegerStream) Average() int {
	return AverageInt(self.Val()...)
}
func (self *IntegerStream) AsFloat32() *Float32Stream {
	result := self.Map2Float32(func(arg int, index int) float32 { return float32(arg) })
	stream := Float32StreamFrom(result)
	return &stream
}
func (self *IntegerStream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg int, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}
func (self *IntegerStream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg int, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}

func (self *IntegerStream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg int, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}
