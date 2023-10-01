package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/utils/strings"
	"reflect"
	"testing"
)

func checkString(s *strings.String, ex string, t *testing.T) bool {
	if s.String() != ex {
		t.Fail()
		fmt.Println("Expected String to be '" + ex + "' but got " + s.String())
		return false
	}
	return true
}

func checkToString(any interface{}, ex string, t *testing.T) bool {
	return _checkToString(any, ex, "xyz", t)
}

func _checkToString(any interface{}, ex, ex2 string, t *testing.T) bool {
	s := strings.StringOf(any)
	_ex := strings.Kind2String(reflect.ValueOf(any)).Add(ex).String()
	_ex2 := strings.Kind2String(reflect.ValueOf(any)).Add(ex2).String()
	if s != _ex && s != _ex2 {
		t.Fail()
		fmt.Println("Expected String to be '" + ex + "' but got '" + s + "'")
		return false
	}
	return true
}

func TestString(t *testing.T) {
	s := strings.New("test")
	if ok := checkString(s, "test", t); !ok {
		return
	}

	s.Add("test")
	if ok := checkString(s, "testtest", t); !ok {
		return
	}

	s.Join(strings.New("test"))
	if ok := checkString(s, "testtesttest", t); !ok {
		return
	}
	if s.IsBlank() {
		t.Failed()
		fmt.Println("Expected s to NOT be blank")
	}
	s = strings.New("")
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
	if ok := _checkToString(map[string]int{"a": 1, "b": 2}, "[a=1,b=2]", "[b=2,a=1]", t); !ok {
		return
	}

	k := reflect.New(reflect.ValueOf("").Type()).Interface()

	if ok := checkToString(k, "", t); !ok {
		return
	}
}
