// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package matcher

import (
	"bytes"

	"github.com/m3db/m3x/pool"

	"github.com/cespare/xxhash"
)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// namespaceRuleSetsMapOptions provides options used when created the map.
type namespaceRuleSetsMapOptions struct {
	InitialSize int
	KeyCopyPool pool.BytesPool
}

// newNamespaceRuleSetsMap returns a new byte keyed map.
func newNamespaceRuleSetsMap(opts namespaceRuleSetsMapOptions) *namespaceRuleSetsMap {
	var (
		copyFn     namespaceRuleSetsMapCopyFn
		finalizeFn namespaceRuleSetsMapFinalizeFn
	)
	if pool := opts.KeyCopyPool; pool == nil {
		copyFn = func(k []byte) []byte {
			return append([]byte(nil), k...)
		}
	} else {
		copyFn = func(k []byte) []byte {
			keyLen := len(k)
			pooled := pool.Get(keyLen)[:keyLen]
			copy(pooled, k)
			return pooled
		}
		finalizeFn = func(k []byte) {
			pool.Put(k)
		}
	}
	return _namespaceRuleSetsMapAlloc(_namespaceRuleSetsMapOptions{
		hash: func(k []byte) namespaceRuleSetsMapHash {
			return namespaceRuleSetsMapHash(xxhash.Sum64(k))
		},
		equals:      bytes.Equal,
		copy:        copyFn,
		finalize:    finalizeFn,
		initialSize: opts.InitialSize,
	})
}
