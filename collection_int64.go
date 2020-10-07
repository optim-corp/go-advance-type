package adtype

import (
	"sort"
)

func Int64StreamOf(arg ...int64) Int64Stream {
	return arg
}
func Int64StreamFrom(arg []int64) Int64Stream {
	return arg
}

func (self *Int64Stream) Add(arg int64) *Int64Stream {
	return self.AddAll(arg)

}

func (self *Int64Stream) AddAll(arg ...int64) *Int64Stream {
	*self = append(*self, arg...)
	return self
}

func (self *Int64Stream) AddSafe(arg *int64) *Int64Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *Int64Stream) AllMatch(fn func(arg int64, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *Int64Stream) AnyMatch(fn func(arg int64, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Int64Stream) Clone() *Int64Stream {
	temp := make([]int64, self.Len())
	copy(temp, *self)
	return (*Int64Stream)(&temp)
}

func (self *Int64Stream) Copy() *Int64Stream {
	return self.Clone()
}

func (self *Int64Stream) Concat(arg []int64) *Int64Stream {
	return self.AddAll(arg...)
}

func (self *Int64Stream) Delete(index int) *Int64Stream {
	return self.DeleteRange(index, index)
}

func (self *Int64Stream) DeleteRange(startIndex int, endIndex int) *Int64Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Int64Stream) Equals(arr []int64) bool {
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
func (self *Int64Stream) Filter(fn func(arg int64, index int) bool) *Int64Stream {
	_array := Int64StreamOf()
	self.ForEach(func(v int64, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *Int64Stream) Find(fn func(arg int64, index int) bool) *int64 {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *Int64Stream) FindIndex(fn func(arg int64, index int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *Int64Stream) First() *int64 {
	return self.Get(0)
}

func (self *Int64Stream) ForEach(fn func(arg int64, index int)) *Int64Stream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *Int64Stream) ForEachRight(fn func(arg int64, index int)) *Int64Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Int64Stream) GroupBy(fn func(arg int64, index int) string) map[string][]int64 {
	m := map[string][]int64{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Int64Stream) GroupByValues(fn func(arg int64, index int) string) [][]int64 {
	tmp := [][]int64{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Int64Stream) IndexOf(arg int64) int {
	for index, _arg := range *self {
		if _arg == arg {
			return index
		}
	}
	return -1
}
func (self *Int64Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Int64Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Int64Stream) Last() *int64 {
	return self.Get(self.Len() - 1)
}

func (self *Int64Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Int64Stream) Limit(limit int) *Int64Stream {
	self.Slice(0, limit)
	return self
}
func (self *Int64Stream) Map(fn func(arg int64, index int) int64) *Int64Stream {
	return self.ForEach(func(v int64, i int) { self.Set(i, fn(v, i)) })
}

func (self *Int64Stream) MapAny(fn func(arg int64, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) Map2Int(fn func(arg int64, index int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) Map2Int32(fn func(arg int64, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) Map2Float32(fn func(arg int64, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) Map2Float64(fn func(arg int64, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) Map2Bool(fn func(arg int64, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) Map2Bytes(fn func(arg int64, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) Map2String(fn func(arg int64, index int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int64Stream) NoneMatch(fn func(arg int64, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *Int64Stream) Get(index int) *int64 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Int64Stream) Peek(fn func(arg *int64, index int)) *Int64Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Int64Stream) Reduce(fn func(result, current int64, index int) int64) *Int64Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Int64Stream) ReduceInit(fn func(result, current int64, index int) int64, initialValue int64) *Int64Stream {
	result := Int64StreamOf()
	self.ForEach(func(v int64, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *Int64Stream) Reverse() *Int64Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *Int64Stream) Replace(fn func(arg int64, index int) int64) *Int64Stream {
	return self.Map(fn)
}

func (self *Int64Stream) Set(index int, val int64) *Int64Stream {
	if len(*self) > index {
		(*self)[index] = val
	}
	return self
}

func (self *Int64Stream) Skip(skip int) *Int64Stream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *Int64Stream) Slice(startIndex int, n int) *Int64Stream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []int64{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *Int64Stream) Sort(fn func(i, j int) bool) *Int64Stream {
	sort.Slice(*self, fn)
	return self
}

func (self *Int64Stream) SortStable(fn func(i, j int) bool) *Int64Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Int64Stream) ToList() []int64 {
	return self.Val()
}

func (self *Int64Stream) Val() []int64 {
	if self == nil {
		return []int64{}
	}
	return *self
}

func (self *Int64Stream) While(fn func(arg int64, index int) bool) *Int64Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Int64Stream) Min() int64 {
	return MinInt64(self.Val()...)
}

func (self *Int64Stream) Max() int64 {
	return MaxInt64(self.Val()...)
}

func (self *Int64Stream) Sum() int64 {
	return SumInt64(self.Val()...)
}

func (self *Int64Stream) Average() int64 {
	return AverageInt64(self.Val()...)
}

func (self *Int64Stream) AsFloat32() *Float32Stream {
	result := self.Map2Float32(func(arg int64, index int) float32 { return float32(arg) })
	stream := Float32StreamFrom(result)
	return &stream
}
func (self *Int64Stream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg int64, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}
func (self *Int64Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg int64, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Int64Stream) AsInt32() *Int32Stream {
	result := self.Map2Int32(func(arg int64, index int) int32 { return int32(arg) })
	stream := Int32StreamFrom(result)
	return &stream
}
