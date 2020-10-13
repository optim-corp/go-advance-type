package adtype

import (
	"math"
	"testing"
)

func TestExchange(t *testing.T) {
	if str := Long2Str(math.MaxInt64); str != "9223372036854775807" {
		t.Fatal("Unexpect Value func Long2Str.", str)
	}
	if i := Str2Int("9223372036854775807", 0); i != 9223372036854775807 {
		t.Fatal("Unexpect Value func Str2Int.", i)
	}
	if i := Str2Int("hogehoge", 999); i != 999 {
		t.Fatal("Unexpect Value func Str2Int.", i)
	}
	if i := Str2Int32(Int2Str(math.MaxInt32), 0); i != 2147483647 {
		t.Fatal("Unexpect Value func Str2Int.", i)
	}
	if i := Str2Int32("9223372036854775807", 0); i != 0 {
		t.Fatal("Unexpect Value func Str2Int32.", i)
	}
	if i := Str2Int32("hogehoge", 999); i != 999 {
		t.Fatal("Unexpect Value func Str2Int32.", i)
	}
	if i := Str2Long("9223372036854775807", 0); i != math.MaxInt64 {
		t.Fatal("Unexpect Value func Str2Long.", i)
	}
	if i := Str2Long("hogehoge", 999); i != 999 {
		t.Fatal("Unexpect Value func Str2Long.", i)
	}
	if i := Str2Long("9223372036854775808", 0); i != 0 {
		t.Fatal("Unexpect Value func Str2Long.", i)
	}
	if i := Str2Double("92233720368547758089999999999999999", 999); i != 92233720368547758089999999999999999 {
		t.Fatal("Unexpect Value func Str2Double.", i)
	}
	if i := Str2Float("92233720368547758089999999999999999", 999); i != 92233720368547758089999999999999999 {
		t.Fatal("Unexpect Value func Str2Double.", i)
	}
	if v := Int2Str(32); v != "32" {
		t.Fatal("Unexpect Value func Int2Str.", v)
	}
	if v := Long2Str(32); v != "32" {
		t.Fatal("Unexpect Value func Long2Str.", v)
	}
	if v := Double2Str(32.5); v != "32.5" {
		t.Fatal("Unexpect Value func Double2Str.", v)
	}
	if v := Float2Str(32); v != "32" {
		t.Fatal("Unexpect Value func Float2Str.", v)
	}
	if v := Any2Int(interface{}(100), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Int.", v)
	}
	if v := Any2Int32(interface{}(int32(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Int32.", v)
	}
	if v := Any2Int32("100", 10); v != 10 {
		t.Fatal("Unexpect Value func Any2Int32.", v)
	}
	if v := Any2Long(interface{}(int64(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Long.", v)
	}
	if v := Any2Float(interface{}(float32(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Float.", v)
	}
	if v := Any2Double(interface{}(float64(100)), 10); v != 100 {
		t.Fatal("Unexpect Value func Any2Double.", v)
	}
	stringStream := StringStreamOf("0", "1", "2", "3", "4", "5")
	intStream := IntegerStreamOf(0, 1, 2, 3, 4, 5)
	int32Stream := Int32StreamOf(0, 1, 2, 3, 4, 5)
	int64Stream := Int64StreamOf(0, 1, 2, 3, 4, 5)
	float32Stream := Float32StreamOf(0, 1, 2, 3, 4, 5)
	float64Stream := Float64StreamOf(0, 1, 2, 3, 4, 5)
	if v := stringStream.Clone().Sum(); v != "012345" {
		t.Fatal("Unexpect Value stringStream Sum.", v)
	}
	if v := intStream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value intStream Sum.", v)
	}
	if v := intStream.Min(); v != 0 {
		t.Fatal("Unexpect Value intStream Min.", v)
	}
	if v := intStream.Max(); v != 5 {
		t.Fatal("Unexpect Value intStream Max.", v)
	}
	if v := intStream.Average(); v != 2 {
		t.Fatal("Unexpect Value intStream Average.", v)
	}
	if v := int32Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value int32Stream Sum.", v)
	}
	if v := int32Stream.Min(); v != 0 {
		t.Fatal("Unexpect Value int32Stream Min.", v)
	}
	if v := int32Stream.Max(); v != 5 {
		t.Fatal("Unexpect Value int32Stream Max.", v)
	}
	if v := int32Stream.Average(); v != 2 {
		t.Fatal("Unexpect Value int32Stream Average.", v)
	}
	if v := int64Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value int64Stream Sum.", v)
	}
	if v := int64Stream.Min(); v != 0 {
		t.Fatal("Unexpect Value int64Stream Min.", v)
	}
	if v := int64Stream.Max(); v != 5 {
		t.Fatal("Unexpect Value int64Stream Max.", v)
	}
	if v := int64Stream.Average(); v != 2 {
		t.Fatal("Unexpect Value int64Stream Average.", v)
	}
	if v := float32Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value float32Stream Sum.", v)
	}
	if v := float32Stream.Min(); v != 0 {
		t.Fatal("Unexpect Value float32Stream Min.", v)
	}
	if v := float32Stream.Max(); v != 5 {
		t.Fatal("Unexpect Value float32Stream Max.", v)
	}
	if v := float32Stream.Average(); v != 2.5 {
		t.Fatal("Unexpect Value float32Stream Average.", v)
	}
	if v := float64Stream.Clone().Sum(); v != 15 {
		t.Fatal("Unexpect Value float64Stream Sum.", v)
	}
}
