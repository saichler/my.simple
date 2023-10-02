package strng

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// Global map that map a type/kind to a method that converts string to that type
var fromstrings = make(map[reflect.Kind]func(string, []reflect.Kind) (reflect.Value, error))

var errorValue = reflect.ValueOf(errors.New("Failed to convert string to instance"))

// initialize the map
func init() {
	fromstrings[reflect.String] = stringFromString
	fromstrings[reflect.Int] = intFromString
	fromstrings[reflect.Int8] = int8FromString
	fromstrings[reflect.Int16] = int16FromString
	fromstrings[reflect.Int32] = int32FromString
	fromstrings[reflect.Int64] = int64FromString
	fromstrings[reflect.Uint] = uintFromString
	fromstrings[reflect.Uint8] = uint8FromString
	fromstrings[reflect.Uint16] = uint16FromString
	fromstrings[reflect.Uint32] = uint32FromString
	fromstrings[reflect.Uint64] = uint64FromString
	fromstrings[reflect.Float32] = float32FromString
	fromstrings[reflect.Float64] = float64FromString
	fromstrings[reflect.Bool] = boolFromString
	fromstrings[reflect.Ptr] = ptrFromString
	fromstrings[reflect.Slice] = sliceFromString
	fromstrings[reflect.Map] = mapFromString
	fromstrings[reflect.Interface] = interfaceFromString
}

// Comvert string to string
func stringFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	return reflect.ValueOf(str), nil
}

// Convert string to int
func intFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	if str == "" {
		return reflect.ValueOf(0), nil
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(i), nil
}

// Convert string to int8
func int8FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(int8(i)), nil
}

// Convert string to int16
func int16FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(int16(i)), nil
}

// Convert string to int32
func int32FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(int32(i)), nil
}

// Convert string to int64
func int64FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(int64(i)), nil
}

// Convert string to uint
func uintFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(uint(i)), nil
}

// Convert string to uint8
func uint8FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	if str == "" {
		return reflect.ValueOf(byte(0)), nil
	}
	return reflect.ValueOf([]byte(str)), nil
}

// Convert string to uint16
func uint16FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(uint16(i)), nil
}

// Convert string to uint32
func uint32FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(uint32(i)), nil
}

// Convert string to uint64
func uint64FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(uint64(i)), nil
}

// Convert string to bool
func boolFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	i, err := strconv.ParseBool(str)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(i), nil
}

// Convert string to float32
func float32FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(float32(f)), nil
}

// Convert string to float64
func float64FromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return errorValue, err
	}
	return reflect.ValueOf(float64(f)), nil
}

// Convert string to pointer
func ptrFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	f := fromstrings[kinds[0]]
	if f != nil {
		v, e := f(str, kinds[1:])
		if e != nil {
			return errorValue, e
		}
		newPtr := reflect.New(v.Type())
		newPtr.Elem().Set(v)
		return newPtr, nil
	}
	return reflect.ValueOf(nil), errors.New("Pointer cloud not be created for kind " + kinds[0].String())
}

// Convert string to interface
func interfaceFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	f := fromstrings[kinds[0]]
	if f != nil {
		v, e := f(str, kinds[1:])
		if e != nil {
			return errorValue, e
		}
		newInt := reflect.New(v.Type())
		newInt.Set(v)
		return newInt, nil
	}
	return reflect.ValueOf(nil), errors.New("Interface cloud not be created for kind " + kinds[0].String())
}

// Convert string to map
func mapFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	str = strings.TrimSpace(str)
	str = str[1 : len(str)-1]
	items := strings.Split(str, ",")
	var newMap *reflect.Value
	for _, item := range items {
		index := strings.Index(item, "=")
		if index == -1 {
			return errorValue, errors.New("one of the items does not contain a = sign")
		}
		keyStr := strings.TrimSpace(item[0:index])
		valueStr := strings.TrimSpace(item[index+1:])
		keyF := fromstrings[kinds[0]]
		valueF := fromstrings[kinds[1]]
		if keyF == nil || valueF == nil {
			return errorValue, errors.New("cannot find either the key type or the value type converter")
		}
		keyV, err := keyF(keyStr, kinds[2:])
		if err != nil {
			return errorValue, err
		}
		valueV, err := valueF(valueStr, kinds[2:])
		if err != nil {
			return errorValue, err
		}
		if newMap == nil {
			m := reflect.MakeMap(reflect.MapOf(keyV.Type(), valueV.Type()))
			newMap = &m
		}
		newMap.SetMapIndex(keyV, valueV)
	}
	return *newMap, nil
}

// Convert string to slice
func sliceFromString(str string, kinds []reflect.Kind) (reflect.Value, error) {
	str = strings.TrimSpace(str)
	// if it is byte array, it will not have square brackets
	if len(str) > 1 && str[0] == '[' {
		str = str[1 : len(str)-1]
	}
	items := strings.Split(str, ",")

	itemF := fromstrings[kinds[0]]
	if itemF == nil {
		return errorValue, errors.New("Cannot find converter item kind " + kinds[0].String())
	}
	defaultValue, err := itemF("", kinds[1:])
	if err != nil {
		return errorValue, err
	}

	//Special case for byte array
	if defaultValue.Kind() == reflect.Uint8 {
		newSlice := reflect.MakeSlice(reflect.SliceOf(defaultValue.Type()), len(str), len(str))
		for i, v := range str {
			newSlice.Index(i).Set(reflect.ValueOf(byte(v)))
		}
		return newSlice, nil
	}

	newSlice := reflect.MakeSlice(reflect.SliceOf(defaultValue.Type()), len(items), len(items))

	for i, item := range items {
		v, e := itemF(item, kinds[1:])
		if e != nil {
			return errorValue, e
		}
		newSlice.Index(i).Set(v)
	}
	return newSlice, nil
}

// Convert string to an instance
func InstanceOf(str string) (interface{}, error) {
	v, e := FromString(str)
	if e != nil {
		return nil, e
	}
	return v.Interface(), e
}

// Conver string to a reflect.value
func FromString(str string) (reflect.Value, error) {
	if str == "" || str == "{0}" {
		return reflect.ValueOf(nil), nil
	}
	v, k, e := parseStringForKinds(str)
	if e != nil {
		return errorValue, e
	}
	f := fromstrings[k[0]]
	if f == nil {
		return errorValue, errors.New("no converted was found for type " + k[0].String())
	}
	return f(v, k[1:])
}

// Extract the kinds from the prefix of the string
func parseStringForKinds(str string) (string, []reflect.Kind, error) {
	if len(str) < 3 {
		return str, nil, errors.New("'" + str + "'lenght is less than 3, which means it is not in the correct format of {kind}...")
	}
	if str[0] != '{' {
		return str, nil, errors.New("'" + str + "' does not start with '{'")
	}
	index := strings.Index(str, "}")
	if index == -1 {
		return str, nil, errors.New("'" + str + "'does not have a closing '}'")
	}
	types := str[1:index]
	result := str[index+1:]
	k, e := parseKinds(types)
	if e != nil {
		return str, k, e
	}
	return result, k, nil
}

// extract the kinds to a list of reflect.Kind
func parseKinds(types string) ([]reflect.Kind, error) {
	split := strings.Split(types, ",")
	kinds := make([]reflect.Kind, len(split))
	for i, v := range split {
		k, e := strconv.Atoi(v)
		if e != nil {
			return nil, e
		}
		kinds[i] = reflect.Kind(k)
	}
	return kinds, nil
}
