package ad_type

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
	temp := Int64StreamOf()
	temp = *self
	return &temp
}

func (self *Int64Stream) Copy() *Int64Stream {
	return self.Clone()
}

func (self *Int64Stream) Concat(arg []int64) *Int64Stream {
	return self.AddAll(arg...)
}

func (self *Int64Stream) Delete(index int) *Int64Stream {
	if len(*self) > index+1 {
		*self = append((*self)[:index], (*self)[index+1:]...)
	} else {
		*self = append((*self)[:index])
	}
	return self
}

func (self *Int64Stream) DeleteRange(startIndex int, endIndex int) *Int64Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}

func (self *Int64Stream) Filter(fn func(arg int64, index int) bool) *Int64Stream {
	_array := []int64{}
	for i, v := range *self {
		if fn(v, i) {
			_array = append(_array, v)
		}
	}
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
	for i, v := range *self {
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
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Int64Stream) ForEachRight(fn func(arg int64, index int)) *Int64Stream {
	for i := len(*self) - 1; i >= 0; i-- {
		fn((*self)[i], i)
	}
	return self
}
func (self *Int64Stream) GroupBy(fn func(arg int64, index int) string) map[string][]int64 {
	m := map[string][]int64{}
	for i, v := range *self {
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
func (self *Int64Stream) IsEmpty() bool {
	if self.Len() == 0 {
		return true
	}
	return false
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

func (self *Int64Stream) Map(fn func(arg int64, index int) int64) *Int64Stream {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	*self = _array
	return self
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

func (self *Int64Stream) Map2Int64(fn func(arg int64, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
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
		return &(*self)[index]
	}
	return nil
}
func (self *Int64Stream) ReduceInit(fn func(result, current int64, index int) int64, initialValue int64) *Int64Stream {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(initialValue, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	*self = result
	return self
}
func (self *Int64Stream) Reduce(fn func(result, current int64, index int) int64) *Int64Stream {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	*self = result
	return self
}
func (self *Int64Stream) ReduceInterface(fn func(result interface{}, current int64, index int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) ReduceString(fn func(result string, current int64, index int) string) []string {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) ReduceInt(fn func(result int, current int64, index int) int) []int {
	result := []int{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) ReduceInt32(fn func(result int32, current int64, index int) int32) []int32 {
	result := []int32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) ReduceInt64(fn func(result int64, current int64, index int) int64) []int64 {
	result := []int64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) ReduceFloat32(fn func(result float32, current int64, index int) float32) []float32 {
	result := []float32{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) ReduceFloat64(fn func(result float64, current int64, index int) float64) []float64 {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) ReduceBool(fn func(result bool, current int64, index int) bool) []bool {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Int64Stream) Reverse() *Int64Stream {
	_array := make([]int64, 0, len(*self))
	for i := len(*self) - 1; i >= 0; i-- {
		_array = append(_array, (*self)[i])
	}
	*self = _array
	return self
}

func (self *Int64Stream) Replace(fn func(arg int64, index int) int64) *Int64Stream {
	for i, v := range *self {
		(*self)[i] = fn(v, i)
	}
	return self
}

func (self *Int64Stream) Set(index int, val int64) {
	if len(*self) > index {
		(*self)[index] = val
	}
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

func (self *Int64Stream) ToList() []int64 {
	return self.Val()
}

func (self *Int64Stream) Val() []int64 {
	if self == nil {
		return []int64{}
	}
	return *self
}
