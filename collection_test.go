package collection_test

import (
	"strconv"
	"testing"

	. "github.com/DanielSvub/collection"
)

func TestDict(t *testing.T) {

	t.Run("basics", func(t *testing.T) {
		d := NewDict[string, int]().
			Set("first", 1).
			Set("second", 2).
			Set("third", 3)

		if d.Get("first") != 1 {
			t.Error("Get should return 1.")
		}
		if !d.Keys().Contains("second") {
			t.Error("Key list should contain the key.")
		}
		if !d.Values().Contains(2) {
			t.Error("Value list should contain the value.")
		}
		if !d.Contains(3) {
			t.Error("Dict should contain value 3.")
		}
		if d.Contains(4) {
			t.Error("Dict should  not contain value 4.")
		}
		if d.Count() != 3 {
			t.Error("Dict should have 3 fields.")
		}
		if d.KeyOf(2) != "second" {
			t.Error("Key for value 2 should be 'second'.")
		}
		if d.Equals(NewDict[string, int]()) {
			t.Error("Dict should not be equal to empty Dict.")
		}
		if d.Pluck("first", "second").Count() != 2 {
			t.Error("Plucked Dict should have 2 fields.")
		}
		d.Unset("third")
		if !NewDict[string, int]().Set("first", 1).Merge(NewDict[string, int]().Set("second", 2)).Equals(d) {
			t.Error("Merge does not work properly.")
		}
		json := d.String()
		if json != `{"first":1,"second":2}` && json != `{"second":2,"first":1}` {
			t.Error("Serialization does not work properly.")
		}
		d.Clear()
		if !d.Empty() {
			t.Error("Dict should be empty.")
		}
	})

	t.Run("equality", func(t *testing.T) {
		if NewDict[string, int]().Set("first", 1).Equals(NewDict[string, int]().Set("second", 2)) {
			t.Error("Equality check does not work properly.")
		}
	})

	t.Run("constructors", func(t *testing.T) {
		if !NewDictFrom(map[string]int{"first": 1, "second": 2}).Equals(NewDict[string, int]().Set("first", 1).Set("second", 2)) {
			t.Error("DictFrom does not work properly.")
		}
	})

	t.Run("export", func(t *testing.T) {
		map1 := map[string]int{"first": 1, "second": 2}
		map2 := NewDict[string, int]().Set("first", 1).Set("second", 2).GoMap()
		for key, value := range map1 {
			if map2[key] != value {
				t.Error("Export to Go Map does not work properly.")
			}
		}
	})

	t.Run("cloning", func(t *testing.T) {
		d := NewDict[string, any]().
			Set("string", "test").
			Set("bool", true).
			Set("int", 1).
			Set("float", 3.14).
			Set("nil", nil)

		if !d.Equals(d.Clone()) {
			t.Error("Dict should be equal to itself.")
		}
	})

	t.Run("functional", func(t *testing.T) {
		d := NewDict[string, int]().
			Set("first", 1).
			Set("second", 2).
			Set("third", 3)
		t1 := NewDict[string, int]()
		d.ForEach(func(key string, value int) { t1.Set(key, value) })
		if !t1.Equals(d) {
			t.Error("ForEach does not work properly.")
		}
		if !d.Map(func(key string, value int) int { return value }).Equals(d) {
			t.Error("Map does not work properly.")
		}
	})

}

func TestList(t *testing.T) {

	t.Run("basics", func(t *testing.T) {
		l := NewList(1, 2, 3)
		if l.Get(0) != 1 {
			t.Error("Get should return 1.")
		}
		if !l.Insert(1, 4).Equals(NewList(1, 4, 2, 3)) {
			t.Error("Element has not been inserted properly.")
		}
		if !l.Replace(1, 5).Equals(NewList(1, 5, 2, 3)) {
			t.Error("Element has not been replaced properly.")
		}
		if !l.Add(6).Delete(1, 4).Equals(NewList(1, 2, 3)) {
			t.Error("Element has not been deleted properly.")
		}
		if l.Insert(3, 4).Pop() != 4 {
			t.Error("Pop does not return a correct value.")
		}
		if !l.Equals(NewList(1, 2, 3)) {
			t.Error("Pop is not working properly.")
		}
		if !l.Clone().Clear().Equals(NewList[int]()) {
			t.Error("List has not been cleared properly.")
		}
		if l.Count() != 3 {
			t.Error("List should have 3 elements.")
		}
		if l.Empty() {
			t.Error("List should not be empty.")
		}
		if !NewList[int]().Empty() {
			t.Error("Empty list should be empty.")
		}
		if !NewList(1, 2).Concat(NewList(3, 4)).Equals(NewList(1, 2, 3, 4)) {
			t.Error("Concatenation does not work properly.")
		}
		if !l.Contains(1) {
			t.Error("List should contain element 1.")
		}
		if l.Contains(4) {
			t.Error("List should not contain element 4.")
		}
		if l.IndexOf(2) != 1 {
			t.Error("Element 2 should be at index 1.")
		}
		if l.IndexOf(4) != -1 {
			t.Error("IndexOf should return -1 if the element is not present.")
		}
		if !l.Clone().Reverse().Equals(NewList(3, 2, 1)) {
			t.Error("Reversing does not work properly.")
		}
	})

	t.Run("equality", func(t *testing.T) {
		if NewList(1).Equals(NewList(2)) {
			t.Error("Equality check does not work properly.")
		}
		if NewList(1).Equals(NewList(1, 2)) {
			t.Error("Equality check does not work properly.")
		}
	})

	t.Run("constructors", func(t *testing.T) {
		if !NewListOf(1, 3).Equals(NewList(1, 1, 1)) {
			t.Error("ListOf does not work properly.")
		}
		if !NewListFrom(make([]int, 3)).Equals(NewList(0, 0, 0)) {
			t.Error("ListFrom does not work properly.")
		}
	})

	t.Run("export", func(t *testing.T) {
		list1 := []int{1, 2}
		list2 := NewList(1, 2).GoSlice()
		for index, value := range list1 {
			if list2[index] != value {
				t.Error("Export to Go Slice does not work properly.")
			}
		}
	})

	t.Run("serialization", func(t *testing.T) {
		if NewList[List[int]](nil).String() != `[null]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList("first", "second").String() != `["first","second"]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList(true, false).String() != `[true,false]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList(1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList(3.14, 5.5).String() != `[3.14,5.5]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList(NewList(1, 2), NewList(3, 4)).String() != `[[1,2],[3,4]]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[any]([]int{1, 2}).String() != `[[1 2]]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[int64](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[int32](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[int16](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[int8](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[uint64](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[uint32](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[uint16](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[uint8](1, 2).String() != `[1,2]` {
			t.Error("Serialization does not work properly.")
		}
		if NewList[float32](3.14, 5.5).String() != `[3.14,5.5]` {
			t.Error("Serialization does not work properly.")
		}
	})

	t.Run("sublist", func(t *testing.T) {
		l := NewList(0, 1, 2, 3, 4)
		if !l.SubList(0, 0).Equals(l) {
			t.Error("SubList(0, 0) should return original list.")
		}
		if !l.SubList(2, 4).Equals(NewList(2, 3)) {
			t.Error("SubList(2, 4) should return two elements.")
		}
		if !l.SubList(0, -2).Equals(NewList(0, 1, 2)) {
			t.Error("SubList(0, -2) should cut last two elements.")
		}
	})

	t.Run("functional", func(t *testing.T) {
		l := NewList(1, 2, 3, 4, 5)
		t1 := NewList[int]()
		l.ForEach(func(value int) { t1.Add(value) })
		if !t1.Equals(l) {
			t.Error("ForEach does not work properly.")
		}
		if !l.Map(func(value int) int { return value }).Equals(l) {
			t.Error("Map does not work properly.")
		}
		if l.Reduce(0, func(sum, x int) int { return sum + x }) != 15 {
			t.Error("Reduce does not work properly.")
		}
		if l.Filter(func(value int) bool { return value <= 3 }).Count() != 3 {
			t.Error("Filter does not work properly.")
		}
	})

	t.Run("numeric", func(t *testing.T) {
		if NewList(2.0, 4.0, 3.0, 5.0, 1.0).Max() != 5.0 {
			t.Error("Float max does not work.")
		}
		if NewList(2, 4, 3, 5, 1).Max() != 5.0 {
			t.Error("Int max does not work.")
		}
		if NewList(2.0, 4.0, 3.0, 5.0, 1.0).Min() != 1.0 {
			t.Error("Float min does not work.")
		}
		if NewList(2, 4, 3, 5, 1).Min() != 1.0 {
			t.Error("Min does not work.")
		}
		if NewList(1.0, 4.0, 5.0).Sum() != 10.0 {
			t.Error("Float sum does not work.")
		}
		if NewList(1, 4, 5).Sum() != 10.0 {
			t.Error("Int sum does not work.")
		}
		if NewList(1.0, 4.0, 5.0).Prod() != 20.0 {
			t.Error("Float prod does not work.")
		}
		if NewList(1, 4, 5).Prod() != 20.0 {
			t.Error("Int prod does not work.")
		}
		if NewList(0.0, 5.0, 5.0, 10.0).Avg() != 5.0 {
			t.Error("Float avg does not work.")
		}
		if NewList(0, 5, 5, 10).Avg() != 5.0 {
			t.Error("Int avg does not work.")
		}
		emptyInt := NewList[int]()
		if emptyInt.Min() != 0 {
			t.Error("Min of empty list does not return 0.")
		}
		if emptyInt.Max() != 0 {
			t.Error("Max of empty list does not return 0.")
		}
		if emptyInt.Sum() != 0 {
			t.Error("Sum of empty list does not return 0.")
		}
		if emptyInt.Prod() != 0 {
			t.Error("Prod of empty list does not return 0.")
		}
		emptyFloat := NewList[float64]()
		if emptyFloat.Min() != 0 {
			t.Error("Min of empty list does not return 0.")
		}
		if emptyFloat.Max() != 0 {
			t.Error("Max of empty list does not return 0.")
		}
		if emptyFloat.Sum() != 0 {
			t.Error("Sum of empty list does not return 0.")
		}
		if emptyFloat.Prod() != 0 {
			t.Error("Prod of empty list does not return 0.")
		}
	})

	t.Run("sorting", func(t *testing.T) {
		if !NewList(2, 4, 3, 5, 1).Sort().Equals(NewList(1, 2, 3, 4, 5)) {
			t.Error("Ascending int sorting does not work properly.")
		}
		if !NewList(2.0, 4.0, 3.0, 5.0, 1.0).Sort().Equals(NewList(1.0, 2.0, 3.0, 4.0, 5.0)) {
			t.Error("Ascending float sorting does not work properly.")
		}
		if !NewList("b", "c", "a").Sort().Equals(NewList("a", "b", "c")) {
			t.Error("Ascending string sorting does not work properly.")
		}
	})

}

func TestTools(t *testing.T) {

	t.Run("mapList", func(t *testing.T) {
		l := NewList(1, 2, 3)
		t1 := NewList("1", "2", "3")
		if !MapList(l, func(value int) string {
			return strconv.Itoa(value)
		}).Equals(t1) {
			t.Error("MapList does not work properly.")
		}
	})

	t.Run("mapDict", func(t *testing.T) {
		o := NewDict[string, int]().
			Set("first", 1).
			Set("second", 2).
			Set("third", 3)
		t1 := NewDict[string, string]().
			Set("first", "1").
			Set("second", "2").
			Set("third", "3")
		if !MapDict(o, func(key string, value int) string {
			return strconv.Itoa(value)
		}).Equals(t1) {
			t.Error("MapDict does not work properly.")
		}
	})
}

func TestPanics(t *testing.T) {

	catch := func(msg string) {
		if r := recover(); r == nil {
			t.Errorf(msg)
		}
	}

	t.Run("uninitDict", func(t *testing.T) {
		defer catch("setting to uninitialized dict did not cause panic")
		var uninit map[string]int
		NewDictFrom(uninit).Set("first", 1)
	})

	t.Run("keyCheck", func(t *testing.T) {
		defer catch("unsetting non-existing key did not cause panic")
		NewDict[string, int]().Unset("test")
	})

	t.Run("valueCheck", func(t *testing.T) {
		defer catch("unsetting non-existing value did not cause panic")
		NewDict[string, int]().KeyOf(1)
	})

	t.Run("uninitList", func(t *testing.T) {
		defer catch("setting to uninitialized list did not cause panic")
		var uninit []int
		NewListFrom(uninit).Add(1)
	})

	t.Run("indexCheck", func(t *testing.T) {
		defer catch("deleting non-existing element did not cause panic")
		NewList[int]().Delete(0)
	})

	t.Run("emptyPop", func(t *testing.T) {
		defer catch("poping from empty list did not cause panic")
		NewList[int]().Pop()
	})

	t.Run("sublist1", func(t *testing.T) {
		defer catch("sublist ending index out of range did not cause panic")
		NewList[int]().SubList(0, 1)
	})

	t.Run("sublist2", func(t *testing.T) {
		defer catch("sublist starting index higher than ending index did not cause panic")
		NewList[int]().SubList(1, 0)
	})

	t.Run("sublist3", func(t *testing.T) {
		defer catch("sublist ending index out of range did not cause panic")
		NewList[int]().SubList(-1, 0)
	})

	t.Run("sort", func(t *testing.T) {
		defer catch("sorting unsortable list did not cause panic")
		NewList[bool]().Sort()
	})

	t.Run("min", func(t *testing.T) {
		defer catch("getting min of non-numeric list did not cause panic")
		NewList[string]().Min()
	})

	t.Run("max", func(t *testing.T) {
		defer catch("getting max of non-numeric list did not cause panic")
		NewList[string]().Max()
	})

	t.Run("sum", func(t *testing.T) {
		defer catch("getting sum of non-numeric list did not cause panic")
		NewList[string]().Sum()
	})

	t.Run("prod", func(t *testing.T) {
		defer catch("getting prod of non-numeric list did not cause panic")
		NewList[string]().Prod()
	})

}
