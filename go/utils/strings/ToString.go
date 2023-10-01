package strings

import (
	"reflect"
	"strconv"
)

// Global map from kind/type into a function that converts this type instance into a string representation
var tostrings = make(map[reflect.Kind]func(reflect.Value) string)

// Initialize, map between the kind and the function that handle it
func init() {
	tostrings[reflect.String] = stringToString
	tostrings[reflect.Int] = intToString
	tostrings[reflect.Int8] = intToString
	tostrings[reflect.Int16] = intToString
	tostrings[reflect.Int32] = intToString
	tostrings[reflect.Int64] = intToString
	tostrings[reflect.Uint] = uintToString
	tostrings[reflect.Uint8] = uintToString
	tostrings[reflect.Uint16] = uintToString
	tostrings[reflect.Uint32] = uintToString
	tostrings[reflect.Uint64] = uintToString
	tostrings[reflect.Float32] = float32ToString
	tostrings[reflect.Float64] = float64ToString
	tostrings[reflect.Bool] = boolToString
	tostrings[reflect.Ptr] = ptrToString
	tostrings[reflect.Slice] = sliceToString
	tostrings[reflect.Map] = mapToString
	tostrings[reflect.Interface] = interfaceToString
	tostrings[reflect.Struct] = structToString
}

// StringOf Accept an instance of any kind and convert it to a String
func StringOf(any interface{}) string {
	val := reflect.ValueOf(any)
	return ToString(val)
}

func ToString(value reflect.Value) string {
	return Kind2String(value).Add(toString(value)).String()
}

// ToString Accepts a value of reflect.value and return its string representation
func toString(value reflect.Value) string {
	if !value.IsValid() {
		return ""
	}
	tostring := tostrings[value.Kind()]
	if tostring == nil {
		panic("No ToString for kind:" + value.Kind().String() + ":" + value.String())
	}
	return tostring(value)
}

// ToString of a String
func stringToString(value reflect.Value) string {
	return value.String()
}

// ToString of an int, int8, int16, int32, int64
func intToString(value reflect.Value) string {
	return strconv.Itoa(int(value.Int()))
}

// ToString of an uint, uint8, uint16, uint32, uint64
func uintToString(value reflect.Value) string {
	return strconv.Itoa(int(value.Uint()))
}

// ToString of a float32
func float32ToString(value reflect.Value) string {
	return strconv.FormatFloat(float64(value.Float()), 'f', -1, 32)
}

// ToString of a float64
func float64ToString(value reflect.Value) string {
	return strconv.FormatFloat(float64(value.Float()), 'f', -1, 64)
}

// ToString of a boolean
func boolToString(value reflect.Value) string {
	if value.Bool() {
		return "true"
	} else {
		return "false"
	}
}

// ToString of a pointer
func ptrToString(value reflect.Value) string {
	if value.IsNil() {
		return "<Nil>"
	}
	return toString(value.Elem())
}

// ToString of a struct
// @TODO - Implement properly
func structToString(value reflect.Value) string {
	return value.String()
}

// ToString of a slice
// format is [<elem>,<elem>,...]
func sliceToString(value reflect.Value) string {
	// Special case if the value is a byte array
	b, ok := value.Interface().([]byte)
	if ok {
		// create a string out of the byte array
		return string(b)
	}

	//If the slice is empty, return empty square brackets
	if value.Len() == 0 {
		return "[]"
	}

	//Return the elements of the slice inside square brackets & delimited by comma
	result := New("[")
	for i := 0; i < value.Len(); i++ {
		if i != 0 {
			result.Add(",")
		}
		elem := value.Index(i)
		result.Add(toString(elem))
	}
	result.Add("]")
	return result.String()
}

// ToStrng of a map
// formst is [<key>=<value,<key>=<value],...]
func mapToString(value reflect.Value) string {
	mapkeys := value.MapKeys()
	if len(mapkeys) == 0 {
		return "[]"
	}
	result := New("[")
	for i, key := range mapkeys {
		if i != 0 {
			result.Add(",")
		}
		val := value.MapIndex(key)
		result.Add(toString(key))
		result.Add("=")
		result.Add(toString(val))
	}
	result.Add("]")
	return result.String()
}

// To String of an interface
func interfaceToString(value reflect.Value) string {
	return toString(value.Elem())
}

// Place the kind value inside curly brackets like {5} == int
func Kind2String(value reflect.Value) *String {
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Ptr {
		s := New("{")
		s.Add(strconv.Itoa(int(value.Kind())))
		s.Add(",")
		s.Add(strconv.Itoa(int(value.Type().Elem().Kind())))
		s.Add("}")
		return s
	} else if value.Kind() == reflect.Map {
		s := New("{")
		s.Add(strconv.Itoa(int(value.Kind())))
		s.Add(",")
		s.Add(strconv.Itoa(int(value.Type().Key().Kind())))
		s.Add(",")
		s.Add(strconv.Itoa(int(value.Type().Elem().Kind())))
		s.Add("}")
		return s
	}
	s := New("{")
	s.Add(strconv.Itoa(int(value.Kind())))
	s.Add("}")
	return s
}
