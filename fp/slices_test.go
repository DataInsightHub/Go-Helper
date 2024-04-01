package fp_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataInsightHub/Go-Helper/fp"
	"github.com/stretchr/testify/assert"
)

func TestDistinctFunction(t *testing.T) {
	t.Run("Should return a new slice with distinct elements", func(t *testing.T) {
		slice := []string{"a", "b", "c", "a", "b", "c"}
		distinctSlice := fp.Distinct(slice)

		assert.Equal(t, 3, len(distinctSlice), "slice did not contain only distinct elements")
	})

	t.Run("Should return a new slice with distinct elements", func(t *testing.T) {

		type testStruct struct {
			ID   string
			Name string
			IDs  []string
		}

		slice := []testStruct{
			{ID: "1", Name: "a", IDs: []string{"1", "2"}},
			{ID: "2", Name: "b"},
			{ID: "1", Name: "a", IDs: []string{"1", "2"}},
		}

		distinctStructs := fp.Distinct(slice)

		assert.Equal(t, 2, len(distinctStructs), "slice did not contain only distinct elements")
	})
}

func TestDistictBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	t.Run("Should return a new slice with distinct elements", func(t *testing.T) {
		slice := []person{
			{name: "a", age: 1},
			{name: "b", age: 2},
			{name: "a", age: 1},
		}

		distinctSlice := fp.DistinctBy(slice, func(p person) string {
			return p.name
		})

		assert.Equal(t, 2, len(distinctSlice), "slice did not contain only distinct elements")
	})
}

func TestForEachParallel(t *testing.T) {
	type person struct {
		name string
		name2 string 
	}

	t.Run("All persons should have the same values within", func(t *testing.T) {
		slice := []person{
			{name: "a"},
			{name: "b"},
			{name: "c"},
		}

		fp.ForEachParallel(slice, func(index int, p person) {
			slice[index].name2 = p.name
		})

		for i := range slice {
			assert.Equal(t, slice[i].name2, slice[i].name)
		}
	})

	t.Run("All persons should have the same values within, with pointers", func(t *testing.T) {
		slice := []*person{
			{name: "a"},
			{name: "b"},
			{name: "c"},
		}

		fp.ForEachParallel(slice, func(_ int, p *person) {
			p.name2 = p.name
		})

		for i := range slice {
			assert.Equal(t, slice[i].name2, slice[i].name)
		}
	})

	t.Run("Limit test", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		const limit = 2

		var active int32
		err := fp.ForEachParallelWithError(slice, func(_ int, _ int) error{
			n := atomic.AddInt32(&active, 1)
			if n > limit {
				return fmt.Errorf("saw %d active goroutines; want ≤ %d", n, limit)
			}
			time.Sleep(1 * time.Microsecond) // Give other goroutines a chance to increment active.
			atomic.AddInt32(&active, -1)
			return nil
		}, fp.WithLimit(limit))

		if err != nil {
			t.Fatal(err)
		}

	})
}

func TestFilter(t *testing.T) {

	t.Run("Should return a new slice with filtered elements", func(t *testing.T) {
		slice := []string{"a", "b", "c", "a", "b", "c"}
		filteredSlice := fp.Filter(slice, func(s string) bool {
			return s == "a"
		})

		assert.Equal(t, 2, len(filteredSlice), "slice did not contain only filtered elements")
	})
}

func TestReduce(t *testing.T) {
	t.Run("Should return a new slice with reduced elements", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		reducedSlice := fp.Reduce(slice, "", func(acc string, s string) string {
			return acc + s
		})

		assert.Equal(t, "abcdef", reducedSlice, "slice did not contain only reduced elements")
	})
}

func TestChunks(t *testing.T) {
	t.Run("Should return a new slice with each element is another slice with the same length", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		chunks := fp.Chunks(slice, 2)

		assert.Equal(t, 3, len(chunks))

		for _, chunk := range chunks {
			assert.Equal(t, 2, len(chunk))
		}
	})
}

func TestHead(t *testing.T) {
	t.Run("Should return the first element of a slice", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		firstElement := fp.Head(slice)

		assert.Equal(t, "a", firstElement, "first element is not correct")
	})

	t.Run("Should return an nul value if slice is empty", func(t *testing.T) {
		var slice []string
		val := fp.Head(slice)

		assert.Equal(t, val, "", "should be empty string")
	})
}

func TestTail(t *testing.T) {
	t.Run("Should return a new slice without the first element", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		tailSlice := fp.Tail(slice)

		assert.Equal(t, 5, len(tailSlice), "slice should not contain the first element")
		assert.Equal(t, "b", tailSlice[0], "first element is not correct")
	})

	t.Run("Should return an empty slice if slice is empty", func(t *testing.T) {
		var slice []string
		tailSlice := fp.Tail(slice)

		assert.Equal(t, 0, len(tailSlice), "slice should be empty")
	})

	t.Run("Should return empty slice if only one element in original slice", func(t *testing.T) {
		slice := []string{"a"}
		tailSlice := fp.Tail(slice)

		assert.Equal(t, 0, len(tailSlice), "slice should have one element")
	})
}

func TestLast(t *testing.T) {
	t.Run("Should return the last element of a slice", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		lastElement := fp.Last(slice)

		assert.Equal(t, "f", lastElement, "last element is not correct")
	})

	t.Run("Should return an nul value if slice is empty", func(t *testing.T) {
		var slice []string
		val := fp.Last(slice)

		assert.Equal(t, val, "", "should be empty string")
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Should return true if slice is empty", func(t *testing.T) {
		var slice []string
		isEmpty := fp.IsEmptySlice(slice)

		assert.True(t, isEmpty, "should be true")
	})

	t.Run("Should return false if slice is not empty", func(t *testing.T) {
		slice := []string{"a"}
		isEmpty := fp.IsEmptySlice(slice)

		assert.False(t, isEmpty, "should be false")
	})
}

func TestCopySlice(t *testing.T) {
	t.Run("Should return a copy of the slice", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		cs := fp.CopySlice(slice)

		assert.Equal(t, slice, cs, "should be equal")
	})

	t.Run("Should return a copy of the slice", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		cs := fp.CopySlice(slice)

		slice[0] = "z"

		assert.NotEqual(t, slice, cs, "should be equal")
		assert.Equal(t, "a", cs[0], "copy should not be changed")
	})
}

func TestConcat(t *testing.T) {
	t.Run("Should return the concatenation of 3 slice", func(t *testing.T) {
		slice1 := []string{"a"}
		slice2 := []string{"b"}
		slice3 := []string{"c"}
		concat := fp.Concat(slice1, slice2, slice3)

		assert.Equal(t, 3, len(concat), "should have the length 3")
		assert.Equal(t, "a", concat[0], "should have 'a' at the index 0")
		assert.Equal(t, "b", concat[1], "should have 'b' at the index 1")
		assert.Equal(t, "c", concat[2], "should have 'c' at the index 2")
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("Should return a new slice with flattened elements", func(t *testing.T) {
		slice := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}}
		flattenedSlice := fp.FlatMap(slice, func(s []string) []string {
			s[0] = "1"
			return s
		})

		assert.Equal(t, 6, len(flattenedSlice), "slice did not contain only flattened elements")
		assert.Equal(t, "1", flattenedSlice[0], "slice did not contain only flattened elements")
		assert.Equal(t, "1", flattenedSlice[5], "slice did not contain only flattened elements")
	})
}

func TestMap(t *testing.T) {
	t.Run("Should return a new slice with mapped elements", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		mappedSlice := fp.Map(slice, func(s string) string {
			return s + "1"
		})

		assert.Equal(t, 6, len(mappedSlice))
		assert.Equal(t, "a1", mappedSlice[0])
		assert.Equal(t, "f1", mappedSlice[5])
	})
}

func TestFindIndices(t *testing.T) {
	t.Run("Should return a slice of indices of elements that match the predicate", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		indices := fp.FindIndices(slice, func(s string) bool {
			return s == "a" || s == "c" || s == "f"
		})

		assert.Equal(t, 3, len(indices))
		assert.Equal(t, 0, indices[0])
		assert.Equal(t, 2, indices[1])
		assert.Equal(t, 5, indices[2])
	})
}

func TestFind(t *testing.T) {
	t.Run("Should return the first element that matches the predicate", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		elements := fp.Find(slice, func(s string) bool {
			return s == "a" || s == "c" || s == "f"
		})

		assert.Equal(t, []string{"a", "c", "f"}, elements)
	})
}

func TestFlatten(t *testing.T) {
	t.Run("Should return a new slice with flattened elements", func(t *testing.T) {
		slice := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}}
		flattenedSlice := fp.Flatten(slice)

		assert.Equal(t, 6, len(flattenedSlice), "slice did not contain only flattened elements")
		assert.Equal(t, "a", flattenedSlice[0], "slice did not contain only flattened elements")
		assert.Equal(t, "f", flattenedSlice[5], "slice did not contain only flattened elements")
	})
}

func TestReplace(t *testing.T) {
	t.Run("Should replace the element at the given index", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		newSlice, replaced := fp.Replace(slice, "z", func(str string) bool {
			return str == "a"
		})

		assert.Equal(t, 1, replaced, "should replace one element")
		assert.Equal(t, "c", newSlice[2], "element was not replaced")
		assert.Equal(t, "z", newSlice[0], "element was replaced")
		assert.Equal(t, "f", newSlice[5], "element was not replaced")
	})
}

func TestForEach(t *testing.T) {
	t.Run("Should iterate over the slice and call the given function", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d", "e", "f"}
		var newSlice []string
		fp.ForEach(slice, func(s string) {
			newSlice = append(newSlice, s)
		})

		assert.Equal(t, slice, newSlice, "slice should be equal")
	})
}

func TestPop(t *testing.T) {

	t.Run("should return last element and remaining slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		popped, remaining, success := fp.Pop(slice)
		assert.Equal(t, 3, popped)
		assert.Equal(t, []int{1, 2}, remaining)
		assert.Equal(t, true, success)
	})

	t.Run("should return last element and empty slice", func(t *testing.T) {
		slice := []int{1}
		popped, remaining, success := fp.Pop(slice)
		assert.Equal(t, 1, popped)
		assert.Equal(t, []int{}, remaining)
		assert.Equal(t, true, success)
	})

	t.Run("should return default element, empty slice and not success", func(t *testing.T) {
		var slice []int
		popped, remaining, success := fp.Pop(slice)
		assert.Equal(t, 0, popped)
		assert.Equal(t, []int{}, remaining)
		assert.Equal(t, false, success)
	})
}

func TestShift(t *testing.T) {

	t.Run("should return first element and remaining slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		popped, remaining, success := fp.Shift(slice)
		assert.Equal(t, 1, popped)
		assert.Equal(t, []int{2, 3}, remaining)
		assert.Equal(t, true, success)
	})

	t.Run("should return first element and empty slice", func(t *testing.T) {
		slice := []int{1}
		popped, remaining, success := fp.Shift(slice)
		assert.Equal(t, 1, popped)
		assert.Equal(t, []int{}, remaining)
		assert.Equal(t, true, success)
	})

	t.Run("should return default element, empty slice and not success", func(t *testing.T) {
		var slice []int
		popped, remaining, success := fp.Shift(slice)
		assert.Equal(t, 0, popped)
		assert.Equal(t, []int{}, remaining)
		assert.Equal(t, false, success)
	})
}

func TestGroupBy(t *testing.T) {

	t.Run("should return a grouped map", func(t *testing.T) {
		slice := []int{1, 1, 2, 2}
		groupedMap := fp.GroupBy(slice, func(value int) int { return value })
		assert.Equal(t, 2, len(groupedMap))
		assert.Equal(t, []int{1, 1}, groupedMap[1])
		assert.Equal(t, []int{2, 2}, groupedMap[2])
	})

}
