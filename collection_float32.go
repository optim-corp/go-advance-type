package adtype

import (
	"sort"
)

func Float32StreamOf(arg ...float32) Float32Stream {
	return arg
}
func Float32StreamFrom(arg []float32) Float32Stream {
	return arg
}

func (self *Float32Stream) Add(arg float32) *Float32Stream {
	return self.AddAll(arg)

}

func (self *Float32Stream) AddAll(arg ...float32) *Float32Stream {
	*self = append(*self, arg...)
	return self
}

func (self *Float32Stream) AddSafe(arg *float32) *Float32Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *Float32Stream) AllMatch(fn func(arg float32, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *Float32Stream) AnyMatch(fn func(arg float32, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Float32Stream) Clone() *Float32Stream {
	temp := make([]float32, self.Len())
	copy(temp, *self)
	return (*Float32Stream)(&temp)
}

func (self *Float32Stream) Copy() *Float32Stream {
	return self.Clone()
}

func (self *Float32Stream) Concat(arg []float32) *Float32Stream {
	return self.AddAll(arg...)
}

func (self *Float32Stream) Delete(index int) *Float32Stream {
	return self.DeleteRange(index, index)
}

func (self *Float32Stream) DeleteRange(startIndex int, endIndex int) *Float32Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Float32Stream) Equals(arr []float32) bool {
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
func (self *Float32Stream) Filter(fn func(arg float32, index int) bool) *Float32Stream {
	_array := Float32StreamOf()
	self.ForEach(func(v float32, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *Float32Stream) Find(fn func(arg float32, index int) bool) *float32 {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *Float32Stream) FindIndex(fn func(arg float32, index int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *Float32Stream) First() *float32 {
	return self.Get(0)
}

func (self *Float32Stream) ForEach(fn func(arg float32, index int)) *Float32Stream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *Float32Stream) ForEachRight(fn func(arg float32, index int)) *Float32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Float32Stream) GroupBy(fn func(arg float32, index int) string) map[string][]float32 {
	m := map[string][]float32{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Float32Stream) GroupByValues(fn func(arg float32, index int) string) [][]float32 {
	tmp := [][]float32{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Float32Stream) IndexOf(arg float32) int {
	for index, _arg := range *self {
		if _arg == arg {
			return index
		}
	}
	return -1
}
func (self *Float32Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Float32Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Float32Stream) Last() *float32 {
	return self.Get(self.Len() - 1)
}

func (self *Float32Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Float32Stream) Limit(limit int) *Float32Stream {
	self.Slice(0, limit)
	return self
}
func (self *Float32Stream) Map(fn func(arg float32, index int) float32) *Float32Stream {
	return self.ForEach(func(v float32, i int) { self.Set(i, fn(v, i)) })
}

func (self *Float32Stream) MapAny(fn func(arg float32, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) Map2Int(fn func(arg float32, index int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) Map2Int32(fn func(arg float32, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) Map2Int64(fn func(arg float32, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) Map2Float64(fn func(arg float32, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) Map2Bool(fn func(arg float32, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) Map2Bytes(fn func(arg float32, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) Map2String(fn func(arg float32, index int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float32Stream) NoneMatch(fn func(arg float32, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *Float32Stream) Get(index int) *float32 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Float32Stream) Peek(fn func(arg *float32, index int)) *Float32Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Float32Stream) Reduce(fn func(result, current float32, index int) float32) *Float32Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Float32Stream) ReduceInit(fn func(result, current float32, index int) float32, initialValue float32) *Float32Stream {
	result := Float32StreamOf()
	self.ForEach(func(v float32, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *Float32Stream) Reverse() *Float32Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *Float32Stream) Replace(fn func(arg float32, index int) float32) *Float32Stream {
	return self.Map(fn)
}

func (self *Float32Stream) Set(index int, val float32) *Float32Stream {
	if len(*self) > index {
		(*self)[index] = val
	}
	return self
}

func (self *Float32Stream) Skip(skip int) *Float32Stream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *Float32Stream) Slice(startIndex int, n int) *Float32Stream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []float32{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *Float32Stream) Sort(fn func(i, j int) bool) *Float32Stream {
	sort.Slice(*self, fn)
	return self
}

func (self *Float32Stream) SortStable(fn func(i, j int) bool) *Float32Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Float32Stream) ToList() []float32 {
	return self.Val()
}

func (self *Float32Stream) Val() []float32 {
	if self == nil {
		return []float32{}
	}
	return *self
}

func (self *Float32Stream) While(fn func(arg float32, index int) bool) *Float32Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Float32Stream) Min() float32 {
	return MinFloat32(self.Val()...)
}

func (self *Float32Stream) Max() float32 {
	return MaxFloat32(self.Val()...)
}

func (self *Float32Stream) Sum() float32 {
	return SumFloat32(self.Val()...)
}

func (self *Float32Stream) Average() float32 {
	return AverageFloat32(self.Val()...)
}

func (self *Float32Stream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg float32, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}

func (self *Float32Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg float32, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Float32Stream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg float32, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}

func (self *Float32Stream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg float32, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}
