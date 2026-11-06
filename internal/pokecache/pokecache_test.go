package pokecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEverything(t *testing.T) {
	a := NewCache(2 * time.Second)
	key := "test"
	val := []byte("test_value")
	a.Add(key, val)
	time.Sleep(1 * time.Second)
	testval, ok := a.Get(key)
	assert.Equal(t, ok, true, "there should be a value")
	assert.Equal(t, testval, val, "value should be equal to \"test\"")
	time.Sleep(2 * time.Second)

	testval, ok = a.Get(key)
	assert.Equal(t, ok, false, "value should not exist anymore")
	assert.NotEqual(t, testval, val, "value should not be equal to \"test\"")
	assert.Nil(t, testval, "value should be nil")
}
