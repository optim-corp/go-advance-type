package ad_type

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
	temp := Float64StreamOf()
	temp = *self
	return &temp
}

func (self *Float64Stream) Copy() *Float64Stream {
	return self.Clone()
}

func (self *Float64Stream) Concat(arg []float64) *Float64Stream {
	return self.AddAll(arg...)
}

func (self *Float64Stream) Delete(index int) *Float64Stream {
	if len(*self) > index+1 {
		*self = append((*self)[:index], (*self)[index+1:]...)
	} else {
		*self = append((*self)[:index])
	}
	return self
}

func (self *Float64Stream) DeleteRange(startIndex int, endIndex int) *Float64Stream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}

func (self *Float64Stream) Filter(fn func(arg float64, index int) bool) *Float64Stream {
	_array := []float64{}
	for i, v := range *self {
		if fn(v, i) {
			_array = append(_array, v)
		}
	}
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
	for i, v := range *self {
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
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *Float64Stream) ForEachRight(fn func(arg float64, index int)) *Float64Stream {
	for i := len(*self) - 1; i >= 0; i-- {
		fn((*self)[i], i)
	}
	return self
}
func (self *Float64Stream) GroupBy(fn func(arg float64, index int) string) map[string][]float64 {
	m := map[string][]float64{}
	for i, v := range *self {
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
func (self *Float64Stream) IsEmpty() bool {
	if self.Len() == 0 {
		return true
	}
	return false
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

func (self *Float64Stream) Map(fn func(arg float64, index int) float64) *Float64Stream {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	*self = _array
	return self
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

func (self *Float64Stream) Map2Float64(fn func(arg float64, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
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
		return &(*self)[index]
	}
	return nil
}
func (self *Float64Stream) ReduceInit(fn func(result, current float64, index int) float64, initialValue float64) *Float64Stream {
	result := []float64{}
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
func (self *Float64Stream) Reduce(fn func(result, current float64, index int) float64) *Float64Stream {
	result := []float64{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	*self = result
	return self
}
func (self *Float64Stream) ReduceInterface(fn func(result interface{}, current float64, index int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn(0.0, v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *Float64Stream) ReduceString(fn func(result string, current float64, index int) string) []string {
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
func (self *Float64Stream) ReduceInt(fn func(result int, current float64, index int) int) []int {
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
func (self *Float64Stream) ReduceInt32(fn func(result int32, current float64, index int) int32) []int32 {
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
func (self *Float64Stream) ReduceInt64(fn func(result int64, current float64, index int) int64) []int64 {
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
func (self *Float64Stream) ReduceFloat32(fn func(result float32, current float64, index int) float32) []float32 {
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
func (self *Float64Stream) ReduceFloat64(fn func(result float64, current float64, index int) float64) []float64 {
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
func (self *Float64Stream) ReduceBool(fn func(result bool, current float64, index int) bool) []bool {
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
func (self *Float64Stream) Reverse() *Float64Stream {
	_array := make([]float64, 0, len(*self))
	for i := len(*self) - 1; i >= 0; i-- {
		_array = append(_array, (*self)[i])
	}
	*self = _array
	return self
}

func (self *Float64Stream) Replace(fn func(arg float64, index int) float64) *Float64Stream {
	for i, v := range *self {
		(*self)[i] = fn(v, i)
	}
	return self
}

func (self *Float64Stream) Set(index int, val float64) {
	if len(*self) > index {
		(*self)[index] = val
	}
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

func (self *Float64Stream) ToList() []float64 {
	return self.Val()
}

func (self *Float64Stream) Val() []float64 {
	if self == nil {
		return []float64{}
	}
	return *self
}