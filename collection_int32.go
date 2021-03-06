package adtype

import (
	"sort"
)

func Int32StreamOf(arg ...int32) Int32Stream {
	return arg
}
func Int32StreamFrom(arg []int32) Int32Stream {
	return arg
}

func (self *Int32Stream) Add(arg int32) *Int32Stream {
	return self.AddAll(arg)

}

func (self *Int32Stream) AddAll(arg ...int32) *Int32Stream {
	*self = append(*self, arg...)
	return self
}

func (self *Int32Stream) AddSafe(arg *int32) *Int32Stream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *Int32Stream) AllMatch(fn func(arg int32, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *Int32Stream) AnyMatch(fn func(arg int32, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *Int32Stream) Clone() *Int32Stream {
	temp := make([]int32, self.Len())
	copy(temp, *self)
	return (*Int32Stream)(&temp)
}

func (self *Int32Stream) Copy() *Int32Stream {
	return self.Clone()
}

func (self *Int32Stream) Concat(arg []int32) *Int32Stream {
	return self.AddAll(arg...)
}

func (self *Int32Stream) Delete(index int) *Int32Stream {
	return self.DeleteRange(index, index)
}

func (self *Int32Stream) DeleteRange(startIndex int, endIndex int) *Int32Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}
func (self *Int32Stream) Equals(arr []int32) bool {
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
func (self *Int32Stream) Filter(fn func(arg int32, index int) bool) *Int32Stream {
	_array := Int32StreamOf()
	self.ForEach(func(v int32, i int) {
		if fn(v, i) {
			_array.Add(v)
		}
	})
	*self = _array
	return self
}

func (self *Int32Stream) Find(fn func(arg int32, index int) bool) *int32 {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *Int32Stream) FindIndex(fn func(arg int32, index int) bool) int {
	for i, v := range self.Val() {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *Int32Stream) First() *int32 {
	return self.Get(0)
}

func (self *Int32Stream) ForEach(fn func(arg int32, index int)) *Int32Stream {
	for i, v := range self.Val() {
		fn(v, i)
	}
	return self
}
func (self *Int32Stream) ForEachRight(fn func(arg int32, index int)) *Int32Stream {
	for i := self.Len() - 1; i >= 0; i-- {
		fn(*self.Get(i), i)
	}
	return self
}
func (self *Int32Stream) GroupBy(fn func(arg int32, index int) string) map[string][]int32 {
	m := map[string][]int32{}
	for i, v := range self.Val() {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *Int32Stream) GroupByValues(fn func(arg int32, index int) string) [][]int32 {
	tmp := [][]int32{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *Int32Stream) IndexOf(arg int32) int {
	for index, _arg := range *self {
		if _arg == arg {
			return index
		}
	}
	return -1
}
func (self *Int32Stream) IsEmpty() bool {
	return self.Len() == 0
}
func (self *Int32Stream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *Int32Stream) Last() *int32 {
	return self.Get(self.Len() - 1)
}

func (self *Int32Stream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}
func (self *Int32Stream) Limit(limit int) *Int32Stream {
	self.Slice(0, limit)
	return self
}
func (self *Int32Stream) Map(fn func(arg int32, index int) int32) *Int32Stream {
	return self.ForEach(func(v int32, i int) { self.Set(i, fn(v, i)) })
}

func (self *Int32Stream) MapAny(fn func(arg int32, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) Map2Int(fn func(arg int32, index int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) Map2Float32(fn func(arg int32, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) Map2Int64(fn func(arg int32, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) Map2Float64(fn func(arg int32, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) Map2Bool(fn func(arg int32, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) Map2Bytes(fn func(arg int32, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) Map2String(fn func(arg int32, index int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *Int32Stream) NoneMatch(fn func(arg int32, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *Int32Stream) Get(index int) *int32 {
	if self.Len() > index && index >= 0 {
		tmp := (*self)[index]
		return &tmp
	}
	return nil
}
func (self *Int32Stream) Peek(fn func(arg *int32, index int)) *Int32Stream {
	for i, v := range *self {
		fn(&v, i)
		self.Set(i, v)
	}
	return self
}
func (self *Int32Stream) Reduce(fn func(result, current int32, index int) int32) *Int32Stream {
	return self.ReduceInit(fn, 0)
}
func (self *Int32Stream) ReduceInit(fn func(result, current int32, index int) int32, initialValue int32) *Int32Stream {
	result := Int32StreamOf()
	self.ForEach(func(v int32, i int) {
		if i == 0 {
			result.Add(fn(initialValue, v, i))
		} else {
			result.Add(fn(result[i-1], v, i))
		}
	})
	*self = result
	return self
}

func (self *Int32Stream) Reverse() *Int32Stream {
	for i, j := 0, self.Len()-1; i < j; i, j = i+1, j-1 {
		(*self)[i], (*self)[j] = (*self)[j], (*self)[i]
	}
	return self
}

func (self *Int32Stream) Replace(fn func(arg int32, index int) int32) *Int32Stream {
	return self.Map(fn)
}

func (self *Int32Stream) Set(index int, val int32) *Int32Stream {
	if len(*self) > index {
		(*self)[index] = val
	}
	return self
}

func (self *Int32Stream) Skip(skip int) *Int32Stream {
	self.Slice(skip, self.Len()-skip)
	return self
}

func (self *Int32Stream) Slice(startIndex int, n int) *Int32Stream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []int32{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *Int32Stream) Sort(fn func(i, j int) bool) *Int32Stream {
	sort.Slice(*self, fn)
	return self
}

func (self *Int32Stream) SortStable(fn func(i, j int) bool) *Int32Stream {
	sort.SliceStable(*self, fn)
	return self
}

func (self *Int32Stream) ToList() []int32 {
	return self.Val()
}

func (self *Int32Stream) Val() []int32 {
	if self == nil {
		return []int32{}
	}
	return *self
}

func (self *Int32Stream) While(fn func(arg int32, index int) bool) *Int32Stream {
	for i, v := range self.Val() {
		if !fn(v, i) {
			break
		}
	}
	return self
}

func (self *Int32Stream) Min() int32 {
	return MinInt32(self.Val()...)
}

func (self *Int32Stream) Max() int32 {
	return MaxInt32(self.Val()...)
}

func (self *Int32Stream) Sum() int32 {
	return SumInt32(self.Val()...)
}

func (self *Int32Stream) Average() int32 {
	return AverageInt32(self.Val()...)
}

func (self *Int32Stream) AsFloat32() *Float32Stream {
	result := self.Map2Float32(func(arg int32, index int) float32 { return float32(arg) })
	stream := Float32StreamFrom(result)
	return &stream
}
func (self *Int32Stream) AsFloat64() *Float64Stream {
	result := self.Map2Float64(func(arg int32, index int) float64 { return float64(arg) })
	stream := Float64StreamFrom(result)
	return &stream
}
func (self *Int32Stream) AsInt() *IntegerStream {
	result := self.Map2Int(func(arg int32, index int) int { return int(arg) })
	stream := IntegerStreamFrom(result)
	return &stream
}
func (self *Int32Stream) AsInt64() *Int64Stream {
	result := self.Map2Int64(func(arg int32, index int) int64 { return int64(arg) })
	stream := Int64StreamFrom(result)
	return &stream
}
func (self *Int32Stream) Contains(arg int32) bool {
	return self.FindIndex(func(_arg int32, index int) bool { return _arg == arg }) != -1
}
func (self *Int32Stream) Distinct() *Int32Stream {
	m := map[int32]bool{}
	self.Filter(func(arg int32, index int) bool {
		_, ok := m[arg]
		if !ok {
			m[arg] = true
		}
		return !ok
	})
	return self
}
