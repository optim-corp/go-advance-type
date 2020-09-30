package ad_type

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
	temp := Int32StreamOf()
	temp = *self
	return &temp
}

func (self *Int32Stream) Copy() *Int32Stream {
	return self.Clone()
}

func (self *Int32Stream) Concat(arg []int32) *Int32Stream {
	return self.AddAll(arg...)
}

func (self *Int32Stream) Delete(index int) *Int32Stream {
	if len(*self) > index+1 {
		*self = append((*self)[:index], (*self)[index+1:]...)
	} else {
		*self = append((*self)[:index])
	}
	return self
}

func (self *Int32Stream) DeleteRange(startIndex int, endIndex int) *Int32Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}

func (self *Int32Stream) Filter(fn func(arg int32, index int) bool) *Int32Stream {
	_array := []int32{}
	for i, v := range *self {
		if fn(v, i) {
			_array = append(_array, v)
		}
	}
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
	for i, v := range *self {
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
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Int32Stream) ForEachRight(fn func(arg int32, index int)) *Int32Stream {
	for i := len(*self) - 1; i >= 0; i-- {
		fn((*self)[i], i)
	}
	return self
}
func (self *Int32Stream) GroupBy(fn func(arg int32, index int) string) map[string][]int32 {
	m := map[string][]int32{}
	for i, v := range *self {
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
func (self *Int32Stream) IsEmpty() bool {
	if self.Len() == 0 {
		return true
	}
	return false
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

func (self *Int32Stream) Map(fn func(arg int32, index int) int32) *Int32Stream {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	*self = _array
	return self
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

func (self *Int32Stream) Map2Int32(fn func(arg int32, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
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

func (self *Int32Stream) Map2Float32(fn func(arg int32, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
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
		return &(*self)[index]
	}
	return nil
}
func (self *Int32Stream) ReduceInit(fn func(result, current int32, index int) int32, initialValue int32) *Int32Stream {
	result := []int32{}
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
func (self *Int32Stream) Reduce(fn func(result, current int32, index int) int32) *Int32Stream {
	result := []int32{}
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
func (self *Int32Stream) ReduceInterface(fn func(result interface{}, current int32, index int) interface{}) []interface{} {
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
func (self *Int32Stream) ReduceString(fn func(result string, current int32, index int) string) []string {
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
func (self *Int32Stream) ReduceInt(fn func(result int, current int32, index int) int) []int {
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
func (self *Int32Stream) ReduceInt32(fn func(result int32, current int32, index int) int32) []int32 {
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
func (self *Int32Stream) ReduceInt64(fn func(result int64, current int32, index int) int64) []int64 {
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
func (self *Int32Stream) ReduceFloat32(fn func(result float32, current int32, index int) float32) []float32 {
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
func (self *Int32Stream) ReduceFloat64(fn func(result float64, current int32, index int) float64) []float64 {
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
func (self *Int32Stream) ReduceBool(fn func(result bool, current int32, index int) bool) []bool {
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
func (self *Int32Stream) Reverse() *Int32Stream {
	_array := make([]int32, 0, len(*self))
	for i := len(*self) - 1; i >= 0; i-- {
		_array = append(_array, (*self)[i])
	}
	*self = _array
	return self
}

func (self *Int32Stream) Replace(fn func(arg int32, index int) int32) *Int32Stream {
	for i, v := range *self {
		(*self)[i] = fn(v, i)
	}
	return self
}

func (self *Int32Stream) Set(index int, val int32) {
	if len(*self) > index {
		(*self)[index] = val
	}
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

func (self *Int32Stream) ToList() []int32 {
	return self.Val()
}

func (self *Int32Stream) Val() []int32 {
	if self == nil {
		return []int32{}
	}
	return *self
}
