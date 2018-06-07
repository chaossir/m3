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

package proto

import (
	"fmt"
	"io"

	"github.com/m3db/m3x/pool"
)

type encoder struct {
	w              io.Writer
	buffer         []byte
	bytesPool      pool.BytesPool
	maxMessageSize int
}

// NewEncoder creates a new encoder, the implementation is not thread safe.
func NewEncoder(w io.Writer, opts BaseOptions) Encoder {
	return newEncoder(w, opts)
}

func newEncoder(w io.Writer, opts BaseOptions) *encoder {
	if opts == nil {
		opts = NewBaseOptions()
	}
	pool := opts.BytesPool()
	return &encoder{
		w:              w,
		buffer:         getByteSliceWithLength(sizeEncodingLength, pool),
		bytesPool:      pool,
		maxMessageSize: opts.MaxMessageSize(),
	}
}

func (e *encoder) Encode(m Marshaler) error {
	size := m.Size()
	if size > e.maxMessageSize {
		return fmt.Errorf("message size %d is larger than maximum supported size %d", size, e.maxMessageSize)
	}
	if err := e.encodeSize(size); err != nil {
		return err
	}
	return e.encodeData(m, size)
}

func (e *encoder) encodeSize(size int) error {
	sizeEncodeDecoder.PutUint32(e.buffer, uint32(size))
	_, err := e.w.Write(e.buffer[:sizeEncodingLength])
	return err
}

func (e *encoder) encodeData(m Marshaler, size int) error {
	e.buffer = growDataBufferIfNeeded(e.buffer, size, e.bytesPool)
	size, err := m.MarshalTo(e.buffer)
	if err != nil {
		return err
	}
	_, err = e.w.Write(e.buffer[:size])
	return err
}

func (e *encoder) resetWriter(w io.Writer) {
	e.w = w
}
