package set_test

import (
	"testing"

	"github.com/cizixs/go-algorithms/set"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	assert.Equal(0, s.Size(), "New created set should be size of 0")
}

func TestAdd(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	data := []interface{}{"string", 32, 3.14, true}
	for i, elem := range data {
		s.Add(elem)
		assert.True(s.Contains(elem), "Set should contain newly added element")
		assert.Equal(i+1, s.Size(), "Set should have %d elements, got %d", i+1, s.Size())
	}
}

func TestContains(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	assert.False(s.Contains("go"), "New created set should contain no element")

	pi := 3.14
	s.Add(pi)
	assert.True(s.Contains(pi), "Set should contain newly added element")
	assert.False(s.Contains(0), "Set should not contain not added element")
}

func TestClear(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	languages := []interface{}{"Python", "C++", "Java", "Golang", "Ruby", "Haskell"}
	s.Add(languages...)

	assert.Equal(len(languages), s.Size(), "Set size should be the elements added")

	s.Clear()
	assert.Equal(0, s.Size(), "Cleared set should be size of 0")
	assert.False(s.Contains("Golang"), "Cleared set should not contain element")
}

func TestCopy(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	languages := []interface{}{"Python", "C++", "Java", "Golang", "Ruby", "Haskell"}
	s.Add(languages...)

	newSet := s.Copy()
	assert.Equal(len(languages), newSet.Size(), "Copied set should have the size of the original set")
	for _, item := range languages {
		assert.True(newSet.Contains(item), "Copied set should contain elements in original set")
	}
}

func TestRemove(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	languages := []interface{}{"Python", "C++", "Java", "Golang", "Ruby", "Haskell"}
	s.Add(languages...)
	s.Remove("Golang")

	assert.False(s.Contains("Golang"), "Removed element should not be in set")
	assert.True(s.Contains("Python"), "Not-Removed element should still be in set")
}

func TestEqual(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	languages := []interface{}{"Python", "C++", "Java", "Golang", "Ruby", "Haskell"}
	s.Add(languages...)

	newSet := s.Copy()

	assert.True(s.Equal(newSet), "Copied set should equal with the original one")
	newSet.Add("Erlang")
	assert.False(s.Equal(newSet), "After adding a new element, set should be different")
	newSet.Remove("C++")
	assert.False(s.Equal(newSet), "Set should be different even they are of the same size")
}

func TestSubsetSuperset(t *testing.T) {
	s := set.New()
	assert := assert.New(t)

	languages := []interface{}{"Python", "C++", "Java", "Golang", "Ruby", "Haskell"}
	s.Add(languages...)

	mastered := set.New()
	mastered.Add("Golang", "Python")

	assert.True(mastered.IsSubset(s))
	assert.True(s.IsSuperset(mastered))

	mastered.Add("Lisp")
	assert.False(mastered.IsSubset(s))
	assert.False(s.IsSuperset(mastered))
}

func TestUnion(t *testing.T) {
	assert := assert.New(t)

	east := set.New()
	east.Add([]interface{}{"China", "Japan", "Korean", "singapore"}...)
	west := set.New()
	west.Add([]interface{}{"England", "America", "France", "German"}...)

	countries := east.Union(west)

	assert.True(countries.Contains("China"))
	assert.True(countries.Contains("France"))
}

func TestIntersection(t *testing.T) {
	assert := assert.New(t)

	phones := set.New("Nexus 5", "iPhone 6S", "iPhone 7", "Mi Phone", "Mate 7")
	macProducts := set.New("iPhone 7", "iPhone 6S", "iPad Air", "Macbook Pro", "iMac")

	iphones := phones.Intersection(macProducts)

	assert.Equal(2, iphones.Size(), "Only two iphones are listed")
	assert.True(iphones.Contains("iPhone 7"))
	assert.True(iphones.Contains("iPhone 6S"))
}

func TestDifference(t *testing.T) {
	assert := assert.New(t)

	primes := set.New(2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37)
	threes := set.New(3, 13, 23, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39)

	primesWithoutThree := primes.Difference(threes)
	assert.True(primesWithoutThree.Contains(5))
	assert.False(primesWithoutThree.Contains(13))
	assert.False(primesWithoutThree.Contains(23))
}

func TestSymmetricDifference(t *testing.T) {
	assert := assert.New(t)

	times3 := set.New(3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39)
	times5 := set.New(5, 10, 15, 20, 25, 30, 35, 40)

	times3Or5Not15 := times3.SymmetricDifference(times5)
	assert.Equal(17, times3Or5Not15.Size())

	assert.True(times3Or5Not15.Contains(3))
	assert.True(times3Or5Not15.Contains(5))
	assert.True(times3Or5Not15.Contains(6))
	assert.True(times3Or5Not15.Contains(9))
	assert.True(times3Or5Not15.Contains(10))
	assert.True(times3Or5Not15.Contains(12))
	assert.False(times3Or5Not15.Contains(15))
	assert.True(times3Or5Not15.Contains(18))
	assert.True(times3Or5Not15.Contains(20))
	assert.True(times3Or5Not15.Contains(21))
	assert.True(times3Or5Not15.Contains(24))
	assert.True(times3Or5Not15.Contains(25))
	assert.True(times3Or5Not15.Contains(27))
	assert.False(times3Or5Not15.Contains(30))
	assert.True(times3Or5Not15.Contains(33))
	assert.True(times3Or5Not15.Contains(35))
	assert.True(times3Or5Not15.Contains(36))
	assert.True(times3Or5Not15.Contains(39))
	assert.True(times3Or5Not15.Contains(40))
}
