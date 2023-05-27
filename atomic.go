// Package buildinfo
package buildinfo

import (
	"sync/atomic"
)

type nocmp [0]func()

// value shadows the type of the same name from sync/atomic
// https://godoc.org/sync/atomic#Value
type value struct {
	_ nocmp // disallow non-atomic comparison

	atomic.Value
}

type atomicString struct {
	_ nocmp // disallow non-atomic comparison

	v value
}

var _zeroString string

// newAtomicString creates a new atomicString.
func newAtomicString(val string) *atomicString {
	x := &atomicString{}
	if val != _zeroString {
		x.store(val)
	}
	return x
}

// load atomically loads the wrapped string.
func (x *atomicString) load() string {
	return unpackString(x.v.Load())
}

// store atomically stores the passed string.
func (x *atomicString) store(val string) {
	x.v.Store(packString(val))
}

func packString(s string) interface{} {
	return s
}

func unpackString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
