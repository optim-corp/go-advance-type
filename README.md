# go-advance-type

痒い所に手が届くユーティリティ

## 型変換
|  変換  |  メソッド  |
| ---- | ---- |
|  int     -> string |  func Int2Str(x int) string  |
|  int64   -> string |  func Long2Str(x int64) string  |
|  float64 -> string |  func Double2Str(x float64) string  |
|  string -> int |  func Str2Int(str string, defaultNumber int) int  |
|  string -> int32 |  func Str2Int32(str string, defaultNumber int32) int32  |
|  string -> int64 |  func Str2Long(str string, defaultNumber int64) int64  |
|  string -> float32 |  func Str2Float(str string, defaultNumber float32) float32   |
|  string -> float64 |  func Str2Double(str string, defaultNumber float64) float64  |


## Max,Min,Sum,Average関数

|  型  |  メソッド  |
| ---- | ---- |
| string  | MaxStr, MinStr, SumStr |
| int     | MaxInt, MinInt, SumInt, AverageInt |
| int32   | MaxInt32, MinInt32, SumInt32, AverageInt32 |
| int64   | MaxInt64, MinInt64, SumInt64, AverageInt64 |
| float32 | MaxFloat32, MinFloat32, SumFloat32, AverageFloat32 |
| float64 | MaxFloat64, MinFloat64, SumFloat64, AverageFloat64 |

## Stream

* string
* float32
* float64
* int
* int32
* int64
* bool

## Usage
 
```go
import adtype "github.com/optim-kazuhiro-seida/go-advance-type"
```
### 型変換

#### From Number

```go
result = adtype.Int2Str(32)
result = adtype.Long2Str(32)
result = adtype.Double2Str(32)
result = adtype.Float2Str(32)
```

#### From String

```go
defaultNumber := 100
result = adtype.Str2Int("100", defaultNumber)
result = adtype.Str2Int32("100", int32(defaultNumber))
result = adtype.Str2Long("100", int64(defaultNumber))
result = adtype.Str2Float("100", float32(defaultNumber))
result = adtype.Str2Double("100", float64(defaultNumber))
```

#### Stream (Array, List, Slice)

##### Stream Usge

###### Add

```go
a := "a"
stream := StringStreamOf("0")     //[0]
stream.Add("1")                   //[0 1]
stream.AddAll("2", "3", "4", "5") //[0 1 2 3 4 5]
stream.Concat([]string{"4", "5"}) //[0 1 2 3 4 5 4 5]
stream.Set(0, "Hogehoge")         //[Hogehoge 0 1 2 3 4 5 4 5]
stream.AddSafe(&a)                //[Hogehoge 0 1 2 3 4 5 4 5 a]
```


###### Delete

```go

stream = StringStreamOf("0", "1", "2", "3", "4", "5", "4", "5") //[0]
stream.DeleteRange(1, 3)                                        //[0 4 5 4 5]
stream.Delete(0)                                                //[4 5 4 5]
stream.Distinct()                                               //[4 5]
stream.Skip(1)                                                  // [5]
stream.Limit(0)                                                 //[]
```


###### Get
```go

stream = StringStreamOf("0", "1", "2", "3", "4", "5", "4", "5")          //[0]
result := stream.Get(-1)                                                 //result = nil
result = stream.Get(2)                                                   //result = *"2"
result = stream.First()                                                  //result = *"0"
result = stream.Last()                                                   //result = *"5"
result = stream.Find(func(arg string, _ int) bool { return arg == "2" }) //result = *"2"
```

###### Match

```go
resultMatch := stream.AllMatch(func(_ string, _ int) bool { return true })       //resultMatch = true
resultMatch = stream.AnyMatch(func(arg string, _ int) bool { return arg == "" }) //resultMatch = false
resultMatch = stream.NoneMatch(func(_ string, _ int) bool { return false })      //resultMatch = true
resultMatch = stream.Contains("5")                                               //resultMatch = true
resultMatch = stream.Equals([]string{"0", "2", "3"})                             //resultMatch = false
resultMatch = stream.IsEmpty()                                                   //resultMatch = true
resultMatch = stream.IsPreset()                                                  //resultMatch = false
stream.FindIndex(func(arg string, _ int) bool { return arg == "3" })             //3
```

###### Array

```go
stream.ForEach(func(arg string, index int) {})                                      //Loop
stream.ForEachRight(func(arg string, index int) {})                                 //Reverse loop
stream.Peek(func(arg *string, index int) {})                                        //Pointer arg
stream.Filter(func(arg string, index int) bool { return index == 2 })               //Filter array
stream.GroupBy(func(arg string, index int) string { return "sample" })              //Make by map
stream.GroupByValue(func(arg string, index int) string { return "sample" })         //Return 2D array
stream.Reduce(func(result, current string, index int) string { return "" })         //Reduce
stream.ReduceInit(func(result, current string, index int) string { return "" }, "") //Reduce and Init value
stream.Slice(0, 5)                                                                  //Trim
stream.Sort(func(i, j int) bool { return true })                                    //sort.Slice
stream.SortStable(func(i, j int) bool { return true })                              //sort.SliceStable
stream.Map(func(arg string, index int) string { return "" })                        //Map2Int, Map2Int32, Map2Bool...
stream.While(func(arg string, index int) bool { return true })                      //While return false
stream.Reverse()                                                                    //Array Reverse
stream.ToList()                                                                     //Return Stream to Array(Type)
stream.Val()                                                                        //Same ToList()
stream.Len()                                                                        //Reutrn Stream length
stream.IndexOf("1")                                                                 //Find index from arg
stream.Clone()                                                                      //Copy Stream
stream.Copy()                                                                       //same Clone()
```

