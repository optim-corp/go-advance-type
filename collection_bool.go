package ad_type

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
	temp := BoolStreamOf()
	temp = *self
	return &temp
}

func (self *BoolStream) Copy() *BoolStream {
	return self.Clone()
}

func (self *BoolStream) Concat(arg []bool) *BoolStream {
	return self.AddAll(arg...)
}

func (self *BoolStream) Delete(index int) *BoolStream {
	if len(*self) > index+1 {
		*self = append((*self)[:index], (*self)[index+1:]...)
	} else {
		*self = append((*self)[:index])
	}
	return self
}

func (self *BoolStream) DeleteRange(startIndex int, endIndex int) *BoolStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}

func (self *BoolStream) Filter(fn func(arg bool, index int) bool) *BoolStream {
	_array := []bool{}
	for i, v := range *self {
		if fn(v, i) {
			_array = append(_array, v)
		}
	}
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
	for i, v := range *self {
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
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *BoolStream) ForEachRight(fn func(arg bool, index int)) *BoolStream {
	for i := len(*self) - 1; i >= 0; i-- {
		fn((*self)[i], i)
	}
	return self
}
func (self *BoolStream) GroupBy(fn func(arg bool, index int) string) map[string][]bool {
	m := map[string][]bool{}
	for i, v := range *self {
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
func (self *BoolStream) IsEmpty() bool {
	if self.Len() == 0 {
		return true
	}
	return false
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

func (self *BoolStream) Map(fn func(arg bool, index int) bool) *BoolStream {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	*self = _array
	return self
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

func (self *BoolStream) Map2Bool(fn func(arg bool, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
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
		return &(*self)[index]
	}
	return nil
}
func (self *BoolStream) ReduceInit(fn func(result, current bool, index int) bool, initialValue bool) *BoolStream {
	result := []bool{}
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
func (self *BoolStream) Reduce(fn func(result, current bool, index int) bool) *BoolStream {
	result := []bool{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	*self = result
	return self
}
func (self *BoolStream) ReduceInterface(fn func(result interface{}, current bool, index int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(false, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *BoolStream) ReduceString(fn func(result string, current bool, index int) string) []string {
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
func (self *BoolStream) ReduceInt(fn func(result int, current bool, index int) int) []int {
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
func (self *BoolStream) ReduceInt32(fn func(result int32, current bool, index int) int32) []int32 {
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
func (self *BoolStream) ReduceInt64(fn func(result int64, current bool, index int) int64) []int64 {
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
func (self *BoolStream) ReduceFloat32(fn func(result float32, current bool, index int) float32) []float32 {
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
func (self *BoolStream) ReduceFloat64(fn func(result float64, current bool, index int) float64) []float64 {
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
func (self *BoolStream) ReduceBool(fn func(result bool, current bool, index int) bool) []bool {
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
func (self *BoolStream) Reverse() *BoolStream {
	_array := make([]bool, 0, len(*self))
	for i := len(*self) - 1; i >= 0; i-- {
		_array = append(_array, (*self)[i])
	}
	*self = _array
	return self
}

func (self *BoolStream) Replace(fn func(arg bool, index int) bool) *BoolStream {
	for i, v := range *self {
		(*self)[i] = fn(v, i)
	}
	return self
}

func (self *BoolStream) Set(index int, val bool) {
	if len(*self) > index {
		(*self)[index] = val
	}
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

func (self *BoolStream) ToList() []bool {
	return self.Val()
}

func (self *BoolStream) Val() []bool {
	if self == nil {
		return []bool{}
	}
	return *self
}
