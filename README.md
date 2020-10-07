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
func main() {
	fmt.Println(adtype.Long2Str(19018019) + "/" + adtype.Int2Str(1029))
	fmt.Println(adtype.Any2Int32("3211", 0))
	fmt.Println(adtype.Any2Int("hogehoge", 0))
	stream := adtype.StringStreamOf("0", "1", "2")
	stream.Add("0").Filter(func(str string, index int) bool {
		return str == "0"
	})
	fmt.Println(stream)
}
```

output

```
19018019/1029
3211
0
[0 0]
```
