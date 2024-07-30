/*
Collection Library for Go
Additional tools
*/
package collection

import (
	"fmt"
	"strconv"
)

/*
Converts a value of any type to string.

Parameters:
  - value - value to convert.

Returns:
  - value converted to string.
*/
func toString(value any) string {
	switch val := any(value).(type) {
	case nil:
		return "null"
	case string:
		return strconv.Quote(val)
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%+v", val)
	}
}

/*
Copies a list and modifies each element by a given mapping function.
The resulting element can be of a different type than the original one.
The function has one parameter, the current element.
The old list remains unchanged.

Parameters:
  - list - old list,
  - function - anonymous function to be executed.

Type parameters:
  - T - type of old list elements,
  - N - type of new list elements.

Returns:
  - new list.
*/
func MapList[T comparable, N comparable](list List[T], function func(x T) N) List[N] {
	new := NewList[N]()
	list.ForEach(func(value T) {
		new.Add(function(value))
	})
	return new
}

/*
Copies a dictionary and modifies each field by a given mapping function.
The resulting element can be of a different type than the original one.
The function has two parameters: key of the current field and its value.
The old dictionary remains unchanged.

Parameters:
  - dict - old dictionary,
  - function - anonymous function to be executed.

Type parameters:
  - K - type of dictionary keys,
  - V - type of old dictionary values,
  - N - type of new dictionary values.

Returns:
  - new dictionary.
*/
func MapDict[K comparable, V comparable, N comparable](dict Dict[K, V], function func(k K, v V) N) Dict[K, N] {
	new := NewDict[K, N]()
	dict.ForEach(func(key K, value V) {
		new.Set(key, function(key, value))
	})
	return new
}
