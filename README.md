# go-advance-type

痒い所に手が届くユーティリティ

```go
func main() {
	fmt.Println(ad_type.Int2Str(1029))
	fmt.Println(ad_type.Long2Str(19018019))
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
1029
19018019
3211
0
[0 0]
```
