package ad_type

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
	temp := IntegerStreamOf()
	temp = *self
	return &temp
}

func (self *IntegerStream) Copy() *IntegerStream {
	return self.Clone()
}

func (self *IntegerStream) Concat(arg []int) *IntegerStream {
	return self.AddAll(arg...)
}

func (self *IntegerStream) Delete(index int) *IntegerStream {
	if len(*self) > index+1 {
		*self = append((*self)[:index], (*self)[index+1:]...)
	} else {
		*self = append((*self)[:index])
	}
	return self
}

func (self *IntegerStream) DeleteRange(startIndex int, endIndex int) *IntegerStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}

func (self *IntegerStream) Filter(fn func(arg int, index int) bool) *IntegerStream {
	_array := []int{}
	for i, v := range *self {
		if fn(v, i) {
			_array = append(_array, v)
		}
	}
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
	for i, v := range *self {
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
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *IntegerStream) ForEachRight(fn func(arg int, index int)) *IntegerStream {
	for i := len(*self) - 1; i >= 0; i-- {
		fn((*self)[i], i)
	}
	return self
}
func (self *IntegerStream) GroupBy(fn func(arg int, index int) string) map[string][]int {
	m := map[string][]int{}
	for i, v := range *self {
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
func (self *IntegerStream) IsEmpty() bool {
	if self.Len() == 0 {
		return true
	}
	return false
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

func (self *IntegerStream) Map(fn func(arg int, index int) int) *IntegerStream {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	*self = _array
	return self
}

func (self *IntegerStream) MapAny(fn func(arg int, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *IntegerStream) Map2Int(fn func(arg int, index int) int) []int {
	_array := make([]int, 0, len(*self))
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

func (self *IntegerStream) Map2Float32(fn func(arg int, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
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
		return &(*self)[index]
	}
	return nil
}
func (self *IntegerStream) ReduceInit(fn func(result, current int, index int) int, initialValue int) *IntegerStream {
	result := []int{}
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
func (self *IntegerStream) Reduce(fn func(result, current int, index int) int) *IntegerStream {
	result := []int{}
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
func (self *IntegerStream) ReduceInterface(fn func(result interface{}, current int, index int) interface{}) []interface{} {
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
func (self *IntegerStream) ReduceString(fn func(result string, current int, index int) string) []string {
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
func (self *IntegerStream) ReduceInt(fn func(result int, current int, index int) int) []int {
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
func (self *IntegerStream) ReduceInt32(fn func(result int32, current int, index int) int32) []int32 {
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
func (self *IntegerStream) ReduceInt64(fn func(result int64, current int, index int) int64) []int64 {
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
func (self *IntegerStream) ReduceFloat32(fn func(result float32, current int, index int) float32) []float32 {
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
func (self *IntegerStream) ReduceFloat64(fn func(result float64, current int, index int) float64) []float64 {
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
func (self *IntegerStream) ReduceBool(fn func(result bool, current int, index int) bool) []bool {
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
func (self *IntegerStream) Reverse() *IntegerStream {
	_array := make([]int, 0, len(*self))
	for i := len(*self) - 1; i >= 0; i-- {
		_array = append(_array, (*self)[i])
	}
	*self = _array
	return self
}

func (self *IntegerStream) Replace(fn func(arg int, index int) int) *IntegerStream {
	for i, v := range *self {
		(*self)[i] = fn(v, i)
	}
	return self
}

func (self *IntegerStream) Set(index int, val int) {
	if len(*self) > index {
		(*self)[index] = val
	}
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

func (self *IntegerStream) ToList() []int {
	return self.Val()
}

func (self *IntegerStream) Val() []int {
	if self == nil {
		return []int{}
	}
	return *self
}
