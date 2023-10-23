package general

import (
	"fmt"
	"reflect"
	"strconv"
)

func StringToInteger(num string) {
	intVar, err := strconv.Atoi(num)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVarBase8, err := strconv.ParseInt(num, 0, 8)
	fmt.Println(intVarBase8, err, reflect.TypeOf(intVarBase8))

	intVarBase16, err := strconv.ParseInt(num, 0, 16)
	fmt.Println(intVarBase16, err, reflect.TypeOf(intVarBase16))

	intVarBase32, err := strconv.ParseInt(num, 0, 32)
	fmt.Println(intVarBase32, err, reflect.TypeOf(intVarBase32))

	intVarBase64, err := strconv.ParseInt(num, 0, 64)
	fmt.Println(intVarBase64, err, reflect.TypeOf(intVarBase64))

}

func StringToFloat() {
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))

	s1 := "-3.141"
	f1, err := strconv.ParseFloat(s1, 8)
	fmt.Println(f1, err, reflect.TypeOf(f1))

	s2 := "A-3141X"
	f2, err := strconv.ParseFloat(s2, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f2, err, reflect.TypeOf(f2))
	}

}

func BoolToString() {
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}

func IntegerToFloat() {
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
	fmt.Println(i32)

	f64 := float64(i32)
	fmt.Println(reflect.TypeOf(f64))
}

func InterFloatConversion() {
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	f64 := float64(f32)
	fmt.Println(reflect.TypeOf(f64))

	f64 = 1097.655698798798
	fmt.Println(f64)

	f32 = float32(f64)
	fmt.Println(f32)
}

func IntToIntConverson() {
	var i int = 10
	fmt.Println(reflect.TypeOf(i))

	i16 := int16(i)
	fmt.Println(reflect.TypeOf(i16))

	i32 := int32(i)
	fmt.Println(reflect.TypeOf(i32))

	i64 := int64(i)
	fmt.Println(reflect.TypeOf(i64))
}

func IntegerToString(num int) {
	// Method 1
	var i int64 = 125
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)

	var s string = strconv.FormatInt(i, 10)
	fmt.Println(reflect.TypeOf(s))

	// Method 2
	b := 1225
	fmt.Println(reflect.TypeOf(b))

	s = fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

func FloatToInteger() {

	// Method 1
	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	var s string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)

	//Method 2
	b := 12.454
	fmt.Println(reflect.TypeOf(b))

	s = fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

}
