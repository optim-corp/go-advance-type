package ad_type

func StringStreamOf(arg ...string) StringStream {
	return arg
}
func StringStreamFrom(arg []string) StringStream {
	return arg
}

func (self *StringStream) Add(arg string) *StringStream {
	return self.AddAll(arg)

}

func (self *StringStream) AddAll(arg ...string) *StringStream {
	*self = append(*self, arg...)
	return self
}

func (self *StringStream) AddSafe(arg *string) *StringStream {
	if arg != nil {
		self.Add(*arg)
	}
	return self

}
func (self *StringStream) AllMatch(fn func(arg string, index int) bool) bool {
	for i, v := range *self {
		if !fn(v, i) {
			return false
		}
	}
	return true
}

func (self *StringStream) AnyMatch(fn func(arg string, index int) bool) bool {
	for i, v := range *self {
		if fn(v, i) {
			return true
		}
	}
	return false
}
func (self *StringStream) Clone() *StringStream {
	temp := StringStreamOf()
	temp = *self
	return &temp
}

func (self *StringStream) Copy() *StringStream {
	return self.Clone()
}

func (self *StringStream) Concat(arg []string) *StringStream {
	return self.AddAll(arg...)
}

func (self *StringStream) Delete(index int) *StringStream {
	if len(*self) > index+1 {
		*self = append((*self)[:index], (*self)[index+1:]...)
	} else {
		*self = append((*self)[:index])
	}
	return self
}

func (self *StringStream) DeleteRange(startIndex int, endIndex int) *StringStream {
	*self = append((*self)[:startIndex], (*self)[endIndex+1:]...)
	return self
}

func (self *StringStream) Filter(fn func(arg string, index int) bool) *StringStream {
	_array := []string{}
	for i, v := range *self {
		if fn(v, i) {
			_array = append(_array, v)
		}
	}
	*self = _array
	return self
}

func (self *StringStream) Find(fn func(arg string, index int) bool) *string {
	i := self.FindIndex(fn)
	if -1 != i {
		return &(*self)[i]
	}
	return nil
}

func (self *StringStream) FindIndex(fn func(arg string, index int) bool) int {
	for i, v := range *self {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

func (self *StringStream) First() *string {
	return self.Get(0)
}

func (self *StringStream) ForEach(fn func(arg string, index int)) *StringStream {
	for i, v := range *self {
		fn(v, i)
	}
	return self
}
func (self *StringStream) ForEachRight(fn func(arg string, index int)) *StringStream {
	for i := len(*self) - 1; i >= 0; i-- {
		fn((*self)[i], i)
	}
	return self
}
func (self *StringStream) GroupBy(fn func(arg string, index int) string) map[string][]string {
	m := map[string][]string{}
	for i, v := range *self {
		key := fn(v, i)
		m[key] = append(m[key], v)
	}
	return m
}
func (self *StringStream) GroupByValues(fn func(arg string, index int) string) [][]string {
	tmp := [][]string{}
	m := self.GroupBy(fn)
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}
func (self *StringStream) IsEmpty() bool {
	if self.Len() == 0 {
		return true
	}
	return false
}
func (self *StringStream) IsPreset() bool {
	return !self.IsEmpty()
}
func (self *StringStream) Last() *string {
	return self.Get(self.Len() - 1)
}

func (self *StringStream) Len() int {
	if self == nil {
		return 0
	}
	return len(*self)
}

func (self *StringStream) Map(fn func(arg string, index int) string) *StringStream {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	*self = _array
	return self
}

func (self *StringStream) MapAny(fn func(arg string, index int) interface{}) []interface{} {
	_array := make([]interface{}, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Int(fn func(arg string, index int) int) []int {
	_array := make([]int, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Int32(fn func(arg string, index int) int32) []int32 {
	_array := make([]int32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Int64(fn func(arg string, index int) int64) []int64 {
	_array := make([]int64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Float32(fn func(arg string, index int) float32) []float32 {
	_array := make([]float32, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Float64(fn func(arg string, index int) float64) []float64 {
	_array := make([]float64, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Bool(fn func(arg string, index int) bool) []bool {
	_array := make([]bool, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2Bytes(fn func(arg string, index int) []byte) [][]byte {
	_array := make([][]byte, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) Map2String(fn func(arg string, index int) string) []string {
	_array := make([]string, 0, len(*self))
	for i, v := range *self {
		_array = append(_array, fn(v, i))
	}
	return _array
}

func (self *StringStream) NoneMatch(fn func(arg string, index int) bool) bool {
	return !self.AnyMatch(fn)
}

func (self *StringStream) Get(index int) *string {
	if self.Len() > index && index >= 0 {
		return &(*self)[index]
	}
	return nil
}
func (self *StringStream) ReduceInit(fn func(result, current string, index int) string, initialValue string) *StringStream {
	result := []string{}
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
func (self *StringStream) Reduce(fn func(result, current string, index int) string) *StringStream {
	result := []string{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	*self = result
	return self
}
func (self *StringStream) ReduceInterface(fn func(result interface{}, current string, index int) interface{}) []interface{} {
	result := []interface{}{}
	for i, v := range *self {
		if i == 0 {
			result = append(result, fn("", v, i))
		} else {
			result = append(result, fn(result[i-1], v, i))
		}
	}
	return result
}
func (self *StringStream) ReduceString(fn func(result string, current string, index int) string) []string {
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
func (self *StringStream) ReduceInt(fn func(result int, current string, index int) int) []int {
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
func (self *StringStream) ReduceInt32(fn func(result int32, current string, index int) int32) []int32 {
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
func (self *StringStream) ReduceInt64(fn func(result int64, current string, index int) int64) []int64 {
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
func (self *StringStream) ReduceFloat32(fn func(result float32, current string, index int) float32) []float32 {
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
func (self *StringStream) ReduceFloat64(fn func(result float64, current string, index int) float64) []float64 {
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
func (self *StringStream) ReduceBool(fn func(result bool, current string, index int) bool) []bool {
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
func (self *StringStream) Reverse() *StringStream {
	_array := make([]string, 0, len(*self))
	for i := len(*self) - 1; i >= 0; i-- {
		_array = append(_array, (*self)[i])
	}
	*self = _array
	return self
}

func (self *StringStream) Replace(fn func(arg string, index int) string) *StringStream {
	for i, v := range *self {
		(*self)[i] = fn(v, i)
	}
	return self
}

func (self *StringStream) Set(index int, val string) {
	if len(*self) > index {
		(*self)[index] = val
	}
}

func (self *StringStream) Slice(startIndex int, n int) *StringStream {
	last := startIndex + n
	if len(*self)-1 < startIndex {
		*self = []string{}
	} else if len(*self) < last {
		*self = (*self)[startIndex:len(*self)]
	} else {
		*self = (*self)[startIndex:last]
	}
	return self
}

func (self *StringStream) ToList() []string {
	return self.Val()
}

func (self *StringStream) Val() []string {
	if self == nil {
		return []string{}
	}
	return *self
}
