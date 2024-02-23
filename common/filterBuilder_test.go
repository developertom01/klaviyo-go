package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	var expectedStr = "filter=equals(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.Equal("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestLessThan(t *testing.T) {
	var expectedStr = "filter=less-than(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.LessThan("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestLessThanOrEqual(t *testing.T) {
	var expectedStr = "filter=less-or-equal(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.LessOrEqual("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestGreaterThan(t *testing.T) {
	var expectedStr = "filter=greater-than(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.GreaterThan("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestGreaterOrEqual(t *testing.T) {
	var expectedStr = "filter=greater-or-equal(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.GreaterOrEqual("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestContains(t *testing.T) {
	var expectedStr = "filter=contains(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.Contains("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestContainsAny(t *testing.T) {
	var expectedStr = "filter=contains-any(field1%2C%5Bvalue1+value2%5D)"
	fb := NewFilterBuilder()
	fb.ContainsAny("field1", []string{"value1", "value2"})

	assert.Equal(t, expectedStr, fb.Build())
}

func TestContainsAll(t *testing.T) {
	var expectedStr = "filter=contains-all(field1%2C%5Bvalue1+value2%5D)"
	fb := NewFilterBuilder()
	fb.ContainsAll("field1", []string{"value1", "value2"})

	assert.Equal(t, expectedStr, fb.Build())
}

func TestEndsWith(t *testing.T) {
	var expectedStr = "filter=ends-with(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.EndsWith("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestStartsWith(t *testing.T) {
	var expectedStr = "filter=starts-with(field1%2Cvalue)"
	fb := NewFilterBuilder()
	fb.StartsWith("field1", "value")

	assert.Equal(t, expectedStr, fb.Build())
}

func TestAny(t *testing.T) {
	var expectedStr = "filter=any(field1%2C%5Bvalue1+value2%5D)"
	fb := NewFilterBuilder()
	fb.Any("field1", []string{"value1", "value2"})

	assert.Equal(t, expectedStr, fb.Build())
}

func TestAnd(t *testing.T) {
	var expectedStr = "filter=and(equals%28field1%252Cvalue1%29%2Cequals%28field2%252Cvalue2%29)"

	fb1 := NewFilterBuilder()
	fb1.Equal("field1", "value1")

	fb2 := NewFilterBuilder()
	fb2.Equal("field2", "value2")

	fb := NewFilterBuilder()
	fb.And(*fb1, *fb2)
	assert.Equal(t, expectedStr, fb.Build())
}

func TestOr(t *testing.T) {
	var expectedStr = "filter=or(equals%28field1%252Cvalue1%29%2Cequals%28field2%252Cvalue2%29)"

	fb1 := NewFilterBuilder()
	fb1.Equal("field1", "value1")

	fb2 := NewFilterBuilder()
	fb2.Equal("field2", "value2")

	fb := NewFilterBuilder()
	fb.Or(*fb1, *fb2)
	assert.Equal(t, expectedStr, fb.Build())
}

func TestNot(t *testing.T) {
	var expectedStr = "filter=not(equals(field1%2Cvalue1))"

	fb1 := NewFilterBuilder()
	fb1.Equal("field1", "value1")

	fb := NewFilterBuilder()
	fb.Not(*fb1)
	assert.Equal(t, expectedStr, fb.Build())
}
