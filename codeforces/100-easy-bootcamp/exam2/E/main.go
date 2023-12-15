package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func isMagicSquare(mat [][]int) bool {
	target := mat[0][0] + mat[0][1] + mat[0][2]

	for i := 0; i < 3; i++ {
		sumByRow := 0
		for j := 0; j < 3; j++ {
			sumByRow += mat[i][j]
		}

		sumByCol := 0
		for j := 0; j < 3; j++ {
			sumByCol += mat[j][i]
		}

		if sumByRow != target || sumByCol != target {
			return false
		}
	}
	sumDiag1 := mat[0][0] + mat[1][1] + mat[2][2]
	sumDiag2 := mat[0][2] + mat[1][1] + mat[2][0]

	return (sumDiag1 == target) && (sumDiag2 == target)
}

func printMat(mat [][]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			writeFormat("%d ", mat[i][j])
		}
		writeLine()
	}
}

func main() {
	defer out.Flush()

	mat := make([][]int, 3)
	for i := 0; i < 3; i++ {
		mat[i] = make([]int, 3)
	}
	mat[0] = readInts()
	mat[1] = readInts()
	mat[2] = readInts()

	for i := 0; i <= 3*int(1e5); i++ {
		tmpMat := make([][]int, 3)
		for i := 0; i < 3; i++ {
			tmpMat[i] = make([]int, 3)
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				tmpMat[i][j] = mat[i][j]
			}
		}

		tmpMat[0][0] = i - (tmpMat[0][1] + tmpMat[0][2])
		tmpMat[1][1] = i - (tmpMat[1][0] + tmpMat[1][2])
		tmpMat[2][2] = i - (tmpMat[2][0] + tmpMat[2][1])

		if tmpMat[0][0] < 0 || tmpMat[1][1] < 0 || tmpMat[2][2] < 0 {
			continue
		}

		if isMagicSquare(tmpMat) {
			printMat(tmpMat)
			return
		}
	}
	printMat(mat)
}

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

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

func write(arg ...interface{})                 { fmt.Fprint(out, arg...) }
func writeLine(arg ...interface{})             { fmt.Fprintln(out, arg...) }
func writeFormat(f string, arg ...interface{}) { fmt.Fprintf(out, f, arg...) }
