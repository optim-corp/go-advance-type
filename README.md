# go-advance-type

痒い所に手が届くユーティリティ

* 型変換
  * 数字から文字列へ変換	
    * int     -> string
    * int32   -> string
    * int64   -> string
    * float32 -> string
    * float64 -> string
  * string 
    * string -> int
    * string -> int32
    * string -> int64
    * string -> float32
    * string -> float64
* Max関数Min関数
  * string
  * int
  * int32
  * int64
  * float32
  * float64
* Stream
  * string
  * float32
  * float64
  * int
  * int32
  * int64
  * bool

```go
func main() {
	fmt.Println(ad_type.Long2Str(19018019) + "/" + ad_type.Int2Str(1029))
	fmt.Println(ad_type.Any2Int32("3211", 0))
	fmt.Println(ad_type.Any2Int("hogehoge", 0))
	stream := ad_type.StringStreamOf("0", "1", "2")
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
