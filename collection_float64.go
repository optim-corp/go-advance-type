package adtype

import (
	"sort"
)

func Float64StreamOf(arg ...float64) Float64Stream {
	return arg
}
func Float64StreamFrom(arg []float64) Float64Stream {
	return arg
}

func (self *Float64Stream) Add(arg float64) *Float64Stream {
	return self.AddAll(arg)

}

func (self *Float64Stream) AddAll(arg ...float64) *Float64Stream {
	*self = append(*self, arg...)
	return self
}

func (self *Float64Stream) AddSafe(arg *float64) *Float64Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *Float64Stream) AllMatch(fn func(arg float64, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *Float64Stream) AnyMatch(fn func(arg float64, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Float64Stream) Clone() *Float64Stream {
	temp := make([]float64, self.Len())
	copy(temp, *self)
	return (*Float64Stream)(&temp)
}

func (self *Float64Stream) Copy() *Float64Stream {
	return self.Clone()
}

func (self *Float64Stream) Concat(arg []float64) *Float64Stream {
	return self.AddAll(arg...)
}

func (self *Float64Stream) Delete(index int) *Float64Stream {
	return self.DeleteRange(index, index)
}

func (self *Float64Stream) DeleteRange(startIndex int, endIndex int) *Float64Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Float64Stream) Equals(arr []float64) bool {
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
func (self *Float64Stream) Filter(fn func(arg float64, index int) bool) *Float64Stream {
	_array := Float64StreamOf()
	self.ForEach(func(v float64, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *Float64Stream) Find(fn func(arg float64, index int) bool) *float64 {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *Float64Stream) FindIndex(fn func(arg float64, index int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *Float64Stream) First() *float64 {
	return self.Get(0)
}

func (self *Float64Stream) ForEach(fn func(arg float64, index int)) *Float64Stream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *Float64Stream) ForEachRight(fn func(arg float64, index int)) *Float64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Float64Stream) GroupBy(fn func(arg float64, index int) string) map[string][]float64 {
	m := map[string][]float64{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Float64Stream) GroupByValues(fn func(arg float64, index int) string) [][]float64 {
	tmp := [][]float64{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Float64Stream) IndexOf(arg float64) int {
	for index, _arg := range *self {
		if _arg == arg {
			return index
		}
	}
	return -1
}
func (self *Float64Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Float64Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Float64Stream) Last() *float64 {
	return self.Get(self.Len() - 1)
}

func (self *Float64Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Float64Stream) Limit(limit int) *Float64Stream {
	self.Slice(0, limit)
	return self
}
func (self *Float64Stream) Map(fn func(arg float64, index int) float64) *Float64Stream {
	return self.ForEach(func(v float64, i int) { self.Set(i, fn(v, i)) })
}

func (self *Float64Stream) MapAny(fn func(arg float64, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2Int(fn func(arg float64, index int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2Int32(fn func(arg float64, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2Int64(fn func(arg float64, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2Float32(fn func(arg float64, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2Bool(fn func(arg float64, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2Bytes(fn func(arg float64, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) Map2String(fn func(arg float64, index int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Float64Stream) NoneMatch(fn func(arg float64, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *Float64Stream) Get(index int) *float64 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Float64Stream) Peek(fn func(arg *float64, index int)) *Float64Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Float64Stream) Reduce(fn func(result, current float64, index int) float64) *Float64Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Float64Stream) ReduceInit(fn func(result, current float64, index int) float64, initialValue float64) *Float64Stream {
	result := Float64StreamOf()
	self.ForEach(func(v float64, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *Float64Stream) Reverse() *Float64Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *Float64Stream) Replace(fn func(arg float64, index int) float64) *Float64Stream {
	return self.Map(fn)
}

func (self *Float64Stream) Set(index int, val float64) *Float64Stream {
	if len(*self) > index {
		(*self)[index] = val
	}
	return self
}

func (self *Float64Stream) Skip(skip int) *Float64Stream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *Float64Stream) Slice(startIndex int, n int) *Float64Stream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []float64{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *Float64Stream) Sort(fn func(i, j int) bool) *Float64Stream {
	sort.Slice(*self, fn)
	return self
}

func (self *Float64Stream) SortStable(fn func(i, j int) bool) *Float64Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Float64Stream) ToList() []float64 {
	return self.Val()
}

func (self *Float64Stream) Val() []float64 {
	if self == nil {
		return []float64{}
	}
	return *self
}

func (self *Float64Stream) While(fn func(arg float64, index int) bool) *Float64Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Float64Stream) Min() float64 {
	return MinFloat64(self.Val()...)
}

func (self *Float64Stream) Max() float64 {
	return MaxFloat64(self.Val()...)
}

func (self *Float64Stream) Sum() float64 {
	return SumFloat64(self.Val()...)
}

func (self *Float64Stream) Average() float64 {
	return AverageFloat64(self.Val()...)
}
func (self *Float64Stream) AsFloat32() *Float32Stream {
	result := self.Map2Float32(func(arg float64, index int) float32 { return float32(arg) })
	stream := Float32StreamFrom(result)
	return &stream
}
func (self *Float64Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg float64, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Float64Stream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg float64, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}

func (self *Float64Stream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg float64, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}
