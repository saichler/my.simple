package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/utils/strng"
	"reflect"
	"strings"
	"testing"
)

var Str = strng.New("")

func init() {
	Str.TypesPrefix = true
}

func checkString(s *strng.String, ex string, t *testing.T) bool {
	if s.String() != ex {
		t.Fail()
		fmt.Println("Expected String to be '" + ex + "' but got " + s.String())
		return false
	}
	return true
}

func checkToString(any interface{}, ex string, t *testing.T) bool {
	return checkToFromString(any, ex, "xyz", t)
}

func checkToFromString(any interface{}, ex, ex2 string, t *testing.T) bool {
	s, e := Str.StringOf(any)
	if e != nil {
		t.Fail()
		fmt.Println("error:", e)
		return false
	}
	fs, e := strng.FromString(s)
	// Until struct is implemented, skip it
	if e != nil && !strings.Contains(s, ",25") {
		t.Fail()
		fmt.Println("error from string:", s, e, fs)
		return false
	}

	_ex := strng.Kind2String(reflect.ValueOf(any)).Add(ex).String()
	_ex2 := strng.Kind2String(reflect.ValueOf(any)).Add(ex2).String()
	if s != _ex && s != _ex2 {
		t.Fail()
		fmt.Println("Expected String to be '" + ex + "' but got '" + s + "'")
		return false
	}
	return true
}

func TestString(t *testing.T) {
	s := strng.New("test")
	if ok := checkString(s, "test", t); !ok {
		return
	}

	s.Add("test")
	if ok := checkString(s, "testtest", t); !ok {
		return
	}

	s.Join(strng.New("test"))
	if ok := checkString(s, "testtesttest", t); !ok {
		return
	}
	if s.IsBlank() {
		t.Failed()
		fmt.Println("Expected s to NOT be blank")
	}
	s = strng.New("")
	if !s.IsBlank() {
		t.Failed()
		fmt.Println("Expected s to be blank")
	}
}

func TestToString(t *testing.T) {
	if ok := checkToString("test", "test", t); !ok {
		return
	}
	if ok := checkToString(int32(4343), "4343", t); !ok {
		return
	}
	if ok := checkToString(uint32(4342), "4342", t); !ok {
		return
	}
	if ok := checkToString(float32(4342.5454), "4342.5454", t); !ok {
		return
	}
	if ok := checkToString(float64(4342.5454), "4342.5454", t); !ok {
		return
	}
	if ok := checkToString(true, "true", t); !ok {
		return
	}
	if ok := checkToString(true, "true", t); !ok {
		return
	}
	if ok := checkToString(nil, "", t); !ok {
		return
	}
	type test struct{}
	if ok := checkToString(&test{}, "<tests.test Value>", t); !ok {
		return
	}
	st := &test{}
	st = nil
	if ok := checkToString(st, "<Nil>", t); !ok {
		return
	}
	if ok := checkToString([]int{}, "[]", t); !ok {
		return
	}
	if ok := checkToString([]int{1, 2, 3}, "[1,2,3]", t); !ok {
		return
	}
	if ok := checkToString([]byte("ABC"), "ABC", t); !ok {
		return
	}
	if ok := checkToFromString(map[string]int{"a": 1, "b": 2}, "[a=1,b=2]", "[b=2,a=1]", t); !ok {
		return
	}

	k := reflect.New(reflect.ValueOf("").Type()).Interface()

	if ok := checkToString(k, "", t); !ok {
		return
	}
}

func TestFromStringPtr(t *testing.T) {
	s, e := strng.InstanceOf("{22,24}test")
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	s1 := *s.(*string)
	if s1 != "test" {
		t.Fail()
		fmt.Println("Expected value to be test but got ", s1)
		return
	}
}

func TestFromStringInt(t *testing.T) {
	v, e := strng.InstanceOf("{2}5")
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	r := v.(int)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int {
		t.Fail()
		fmt.Println("From string failed for int")
		return
	}
}

func TestFromStringInt8(t *testing.T) {
	v, e := strng.InstanceOf("{3}5")
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	r := v.(int8)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int8 {
		t.Fail()
		fmt.Println("From string failed for int8")
		return
	}
}

func TestFromStringInt16(t *testing.T) {
	v, e := strng.InstanceOf("{4}5")
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	r := v.(int16)
	if r != 5 || reflect.ValueOf(r).Kind() != reflect.Int16 {
		t.Fail()
		fmt.Println("From string failed for int16")
		return
	}
}

func TestFromStringFloat32(t *testing.T) {
	v, e := strng.InstanceOf("{13}5.8")
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	r := v.(float32)
	if r != 5.8 || reflect.ValueOf(r).Kind() != reflect.Float32 {
		t.Fail()
		fmt.Println("From string failed for float32")
		return
	}
}

func TestFromStringSlice(t *testing.T) {
	s, e := strng.InstanceOf("{23,24}[a,b]")
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	s1 := s.([]string)
	if s1[0] != "a" {
		t.Fail()
		fmt.Println("value for index 0 was not equale to a")
		return
	}
	if s1[1] != "b" {
		t.Fail()
		fmt.Println("value for index 0 was not equale to b")
		return
	}
}

func TestFromStringMap(t *testing.T) {
	s, e := strng.InstanceOf("{21,24,2}[a=1,b=2]")
	if e != nil {
		t.Fail()
		fmt.Println(e)
		return
	}
	s1 := s.(map[string]int)
	if s1["a"] != 1 {
		t.Fail()
		fmt.Println("value for key 'a' was not found or not equale to 1")
		return
	}
	if s1["b"] != 2 {
		t.Fail()
		fmt.Println("value for key 'b' was not found or not equale to 2")
		return
	}
}
