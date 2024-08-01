# Collection

Collection is a Go library providing generic lists and dictionaries with API inspired by Java collections. It is largely similar to the dynamic [AnyType](https://github.com/DanielSvub/anytype) library.

## Dictionaries

Dictionary is an unordered set of key-value pairs. It is a generic interface with two type parameters: type of keys (K) and type of values (V), which both have to satisfy the comparable constraint. The library provides a default implementation based on built-in Go maps. It is possible to make custom implementations by implementing the `Dict` interface.

### Constructors

- `NewDict[K, V]() Dict[K, V]` - creates a new empty dictionary,
```go
dict := collection.NewDict[string, int]()
```

- `NewDictFrom[K, V](goMap map[K]V) Dict[K, V]` - creates a list from a given Go map.
```go
dict := collection.NewDictFrom(map[string]int{
	"first": 1,
	"second": 2,
    "third": 3,
})
```

### Manipulation With Fields
- `Set(key K, value V) Dict[K, V]` - new value is set as key-value pair,
```go
dict.Set("first", 1)
```

- `Unset(keys ...K) Dict[K, V]` - removes the given keys from the dictionary,
```go
dict.Unset("first", "second")
```

- `Clear() Dict[K, V]` - removes all keys in the dictionary,
```go
dict.Clear()
```

- `Get(key K) V` - acquires a value of a field.
```go
value := dict.Get("first")
```

### Export
- `String() string` - exports the dictionary into a string representation. As long as only JSON supported types are used (strings, numbers, bools, nils, nested dictionaries with string keys and nested lists), the output is a valid JSON,
```go
fmt.Println(dict.String())
```

- `GoMap() map[K]V` - exports the dictionary into a Go map,
```go
var goMap map[string]int
goMap = dict.GoMap()
```

- `Keys() List[K]` - exports all keys of the dictionary into a list,
```go
var keys collection.List
keys = dict.Keys()
```

- `Values() List[V]` - exports all values of the dictionary into a list.
```go
var values collection.List
values = dict.Values()
```

### Features Over Whole Dictionary
- `Clone() Dict[K, V]` - performs a copy of the dictionary. Nested lists and dictionaries are copied by reference,
```go
copy := dict.Clone()
```

- `Count() int` - returns a number of fileds in the dictionary,
```go
for i := 0; i < dict.Count(); i++ {
    // ...
}
```

- `Empty() bool` - checks whether the dictionary is empty,
```go
if dict.Empty() {
    // ...
}
```

- `Equals(another Dict[K, V]) bool` - checks whether all fields of the dictionary are equal to the fields of another dictionary,
```go
if dict.Equals(another) {
    // ...
}
```

- `Merge(another Dict[K, V]) Dict[K, V]` - merges two dictionaries together,
```go
merged := dict.Merge(another)
```

- `Pluck(keys ...K) Dict[K, V]` - creates a new dictionary containing only the selected keys from existing dictionary,
```go
plucked := dict.Pluck("first", "second")
```

- `Contains(value V) bool` - checks whether the dictionary contains a certain value,
```go
if dict.Contains(1) {
    // ...
}
```

- `KeyOf(value V) K` - returns any key containing the given value. It panics if the dictionary does not contain the value,
```go
first := dict.KeyOf(1)
```

- `KeyExists(key K) bool` - checks whether a key exists within the dictionary.
```go
if dict.KeyExists("first") {
    // ...
}
```

### Functional Programming
- `ForEach(function func(K, V)) Dict[K, V]` - executes a given function over an every field of the dictionary,
```go
dict.ForEach(func(key string, value int) {
    // ...
})
```

- `Map(function func(K, V) V) Dict[K, V]` - returns a new dictionary with fields modified by a given function. As methods in Go cannot be generic, the target type has to be the same as the source type. If a type change is needed, check the [Additional Tools](#additional-tools) section.
```go
mapped := dict.Map(func(key string, value int) int {
    // ...
	return newValue
})
```

## Lists

List is an ordered sequence of elements. It is a generic interface with one type parameter: type of elements (T), which has to satisfy the comparable constraint. The library provides a default implementation based on built-in Go slices. It is possible to make custom implementations by implementing the `List` interface.

### Constructors

- `NewList[T](values ...T) List[T]` - initial list elements could be given as variadic arguments,
```go
emptyList := collection.NewList[int]()
list := collection.NewList(1, 2, 3)
```

- `NewListOf[T](value T, count int) List[T]` - creates a list of n repeated values,
```go
list := collection.NewListOf(1, 10)
```

- `NewListFrom[T](slice []T) List[T]` - creates a list from a given Go slice.
```go
list := collection.NewListFrom([]int{1, 2, 3})
```

### Manipulation With Elements
- `Add(val ...T) List[T]` - adds any amount of new elements to the list,
```go
list.Add(1, 2, 3)
```

- `Insert(index int, value T) List[T]` - inserts a new element to a specific position in the list,
```go
list.Insert(1, 2)
```

- `Replace(index int, value T) List[T]` - replaces an existing element,
```go
list.Replace(1, 2)
```

- `Delete(index ...int) List[T]` - removes specified elements,
```go
list.Delete(1, 2)
```

- `Pop() T` - removes the last element from the list and returns it,
```go
last := list.Pop()
```

- `Clear() List[T]` - removes all elements in the list.
```go
list.Clear()
```

- `Get(index int) T` - acquires a value of an element.
```go
value := list.Get(1)
```

### Export
- `String() string` - exports the list into a string representation. As long as only JSON supported types are used (strings, numbers, bools, nils, nested dictionaries with string keys and nested lists), the output is a valid JSON,
```go
fmt.Println(list.String())
```

- `Slice() []T` - exports the list into a Go slice.
```go
var slice []int
slice = list.Slice()
```

### Features Over Whole List
- `Clone() List[T]` - performs a copy of the list. Nested lists and dictionaries are copied by reference,
```go
copy := list.Clone()
```

- `Count() int` - returns a number of elements in the list,
```go
for i := 0; i < list.Count(); i++ {
    // ...
}
```

- `Empty() bool` - checks whether the list is empty,
```go
if list.Empty() {
    // ...
}
```

- `Equals(another List[T]) bool` - checks whether all elements of the list are equal to the elements of another list,
```go
if list.Equals(another) {
    // ...
}
```

- `Concat(another List[T]) List[T]` - concates two lists together,
```go
concated := list.Concat(another)
```

- `SubList(start int, end int) List[T]` - cuts a part of the list,
```go
subList := list.SubList(1, 3)
```

- `Contains(elem T) bool` - checks whether the list contains a certain value,
```go
if list.Contains(1) {
    // ...
}
```

- `IndexOf(elem T) int` - returns a position of the first occurrence of the given value,
```go
index := list.IndexOf(1)
```

- `Sort() List[T]` - sorts the elements in the list. The list has to be either of type string, int or float64,
```go
list.Sort()
```

- `Reverse() List[T]` - reverses the list.
```go
list.Reverse()
```

### Functional Programming
- `ForEach(function func(T)) List[T]` - executes a given function over an every element of the list,
```go
list.ForEach(func(value int) {
    // ...
})
```

- `Map(function func(T) T) List[T]` - returns a new list with elements modified by a given function. As methods in Go cannot be generic, the target type has to be the same as the source type. If a type change is needed, check the [Additional Tools](#additional-tools) section,
```go
mapped := list.Map(func(value int) int {
    // ...
	return newValue
})
```

- `Reduce(initial T, function func(T, T) T) T` - reduces all elements in the list into a single value,
```go
result := list.Reduce(0, func(sum, value int) int {
	return sum + value
})
```

- `Filter(function func(T) bool) List[T]` - filters elements in the list based on a condition.
```go
filtered := list.Filter(func(value int) bool {
    // ...
	return condition
})
```

### Numeric Operations

- `Sum() float64` - computes a sum of all elements in the list. List has to be either of type int or float64,
```go
sum := list.Sum()
```

- `Prod() float64` - computes a product of all elements in the list. List has to be either of type int or float64,
```go
product := list.Prod()
```

- `Avg() float64` - computes an arithmetic mean of all elements in the list. List has to be either of type int or float64,
```go
average := list.Avg()
```

- `Min() float64` - returns a minimum value in the list. List has to be either of type int or float64,
```go
minimum := list.Min()
```

- `Max() float64` - returns a maximum value in the list. List has to be either of type int or float64.
```go
maximum := list.Max()
```

## Additional tools

Because the mapping methods of both dictionary and list always keep types, additional mapping functions are available:

`MapDict[K, V, N](dict Dict[K, V], function func(K, V) N) Dict[K, N]` - returns a new dictionary with fields of an existing dictionary modified by a given function.
```go
mapped := MapDict(func(key string, value int) string {
    // ...
	return newValue
})
```

`MapList[T, N](list List[T], function func(T) N) List[N]` - returns a new list with elements of an existing list modified by a given function.
```go
mapped := MapList(list, func(value int) string {
    // ...
	return newValue
})
```
