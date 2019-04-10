package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChopUpWithEmptyString(t *testing.T) {
	teststr := ""
	count := 3
	m := ChopUp(teststr, count)
	assert.Nil(t, m)
}

func TestChopUpWithLargeCount(t *testing.T) {
	teststr := "abc"
	count := 9
	m := ChopUp(teststr, count)
	assert.Nil(t, m)
}

func TestChopUpWithZeroCount(t *testing.T) {
	teststr := "abc"
	count := 0
	m := ChopUp(teststr, count)
	assert.Nil(t, m)
}

func TestChopUpWithNegativeCount(t *testing.T) {
	teststr := "abc"
	count := -42
	m := ChopUp(teststr, count)
	assert.Nil(t, m)
}

func TestChopUpWithUniqueString(t *testing.T) {
	teststr := "aaaaa"
	count := 3
	m := ChopUp(teststr, count)
	assert.NotNil(t, m)
	assert.Equal(t, 1, len(m))

	for key, _ := range m {
		assert.Equal(t, "aaa", key)
	}
	assert.Equal(t, 3, m["aaa"])
}

func TestChopUpWithThreeResultsString(t *testing.T) {
	teststr := "abcde"
	count := 3
	m := ChopUp(teststr, count)
	assert.Equal(t, 3, len(m))

	referenceValues := []string{
		"abc",
		"bcd",
		"cde",
	}

	assert.True(t, validateKeys(referenceValues, m))
}

func TestChopUpWithRepeatString(t *testing.T) {
	teststr := "abcdeabcde"
	count := 3
	m := ChopUp(teststr, count)
	assert.Equal(t, 5, len(m))

	referenceValues := []string{
		"abc",
		"bcd",
		"cde",
		"dea",
		"eab",
	}

	assert.True(t, validateKeys(referenceValues, m))
}

func validateKeys(ref []string, m map[string]int) bool {
	r := make(map[string]bool)
	for _, val := range ref {
		r[val] = true
	}

	for key, _ := range m {
		if false == r[key] {
			return false
		}
		delete(r, key)
	}
	return true
}
