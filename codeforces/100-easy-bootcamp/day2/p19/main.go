package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var n int
var arr []int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getDivisors(n int) []int {
	var res []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				res = append(res, i)
			} else {
				res = append(res, i, n/i)
			}
		}
	}

	return res
}

func isChallenging(h []int, l, r int) bool {

	for i := l; i <= r-1; i++ {
		if i == 0 || i == n-1 {
			continue
		}

		if h[i] > h[i-1] && h[i] > h[i+1] {
			return true
		}
	}

	return false
}

func main() {
	n = readInts()[0]
	arr = readInts()

	divs := getDivisors(n)

	ans := 0
	for _, div := range divs {
		if div == 1 {
			continue
		}
		isSuccess := true
		for i := 0; i < n; i += div {
			isSuccess = isSuccess && isChallenging(arr, i, i+div)
		}
		if isSuccess {
			ans = max(ans, n/div)
		}
	}

	writeLine(ans)
}

var in = bufio.NewReader(os.Stdin)

type float = float32
type double = float64
type long = int64
type ulong = uint64

func typeOf(arg interface{}) string {
	switch reflect.TypeOf(arg).Kind() {
	case reflect.Int:
		return "int"
	case reflect.Int8:
		return "int8"
	case reflect.Int16:
		return "int16"
	case reflect.Int32:
		return "int32"
	case reflect.Int64:
		return "int64"
	case reflect.Uint:
		return "uint"
	case reflect.Uint8:
		return "uint8"
	case reflect.Uint16:
		return "uint16"
	case reflect.Uint32:
		return "uint32"
	case reflect.Uint64:
		return "uint64"
	case reflect.Float32:
		return "float32"
	case reflect.Float64:
		return "float64"
	case reflect.Complex64:
		return "complex64"
	case reflect.Complex128:
		return "complex128"
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "string"
	case reflect.Chan:
		return "chan"
	case reflect.Ptr:
		return "ptr"
	case reflect.Uintptr:
		return "uintptr"
	case reflect.UnsafePointer:
		return "unsafeptr"
	case reflect.Interface:
		return "interface"
	case reflect.Array:
		return "array"
	case reflect.Slice:
		return "slice"
	case reflect.Map:
		return "map"
	case reflect.Struct:
		return "struct"
	case reflect.Func:
		return "func"
	case reflect.Invalid:
		return "invalid"
	default:
		return "nil"
	}
}

func _scln() string {
	ln, _ := in.ReadString('\n')
	return strings.Trim(ln, " \r\n")
}
func _sc(s []string) string {
	if len(s) == 0 {
		return _scln()
	}
	return s[0]
}

func readLine() string              { return _scln() }
func readString() string            { return _scln() }
func readStrings() []string         { return strings.Split(_scln(), " ") }
func readBool(s ...string) bool     { t, _ := strconv.ParseBool(_sc(s)); return t }
func readByte(s ...string) byte     { t, _ := strconv.ParseUint(_sc(s), 10, 8); return byte(t) }
func readDouble(s ...string) double { t, _ := strconv.ParseFloat(_sc(s), 64); return t }
func readFloat(s ...string) float   { t, _ := strconv.ParseFloat(_sc(s), 32); return float32(t) }
func readInt(s ...string) int       { t, _ := strconv.Atoi(_sc(s)); return t }
func readLong(s ...string) long     { t, _ := strconv.ParseInt(_sc(s), 10, 64); return t }
func readULong(s ...string) ulong   { t, _ := strconv.ParseUint(_sc(s), 10, 64); return t }

func readBools() []bool {
	strs := readStrings()
	arr := make([]bool, len(strs))
	for i, s := range strs {
		arr[i] = readBool(s)
	}
	return arr
}
func readBytes() []byte {
	strs := readStrings()
	arr := make([]byte, len(strs))
	for i, s := range strs {
		arr[i] = readByte(s)
	}
	return arr
}
func readDoubles() []double {
	strs := readStrings()
	arr := make([]double, len(strs))
	for i, s := range strs {
		arr[i] = readDouble(s)
	}
	return arr
}
func readFloats() []float {
	strs := readStrings()
	arr := make([]float, len(strs))
	for i, s := range strs {
		arr[i] = readFloat(s)
	}
	return arr
}
func readInts() []int {
	strs := readStrings()
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i] = readInt(s)
	}
	return arr
}
func readLongs() []long {
	strs := readStrings()
	arr := make([]long, len(strs))
	for i, s := range strs {
		arr[i] = readLong(s)
	}
	return arr
}
func readULongs() []ulong {
	strs := readStrings()
	arr := make([]ulong, len(strs))
	for i, s := range strs {
		arr[i] = readULong(s)
	}
	return arr
}

func write(arg ...interface{})                 { fmt.Print(arg...) }
func writeLine(arg ...interface{})             { fmt.Println(arg...) }
func writeFormat(f string, arg ...interface{}) { fmt.Printf(f, arg...) }
