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

package buffer

import (
	"time"

	"github.com/m3db/m3x/instrument"
)

// OnFullStrategy defines the buffer behavior when the buffer is full.
type OnFullStrategy string

const (
	// ReturnError means an error will be returned
	// on new buffer requests when the buffer is full.
	ReturnError OnFullStrategy = "returnError"

	// DropEarliest means the earlist data in the buffer
	// will be dropped to make room for new buffer requests
	// when the buffer is full.
	DropEarliest OnFullStrategy = "dropEarliest"
)

// Options configs the buffer.
type Options interface {
	// OnFullStrategy returns the strategy when buffer is full.
	OnFullStrategy() OnFullStrategy

	// SetOnFullStrategy sets the strategy when buffer is full.
	SetOnFullStrategy(value OnFullStrategy) Options

	// MaxBufferSize returns the max buffer size.
	MaxBufferSize() int

	// SetMaxBufferSize sets the max buffer size.
	SetMaxBufferSize(value int) Options

	// CleanupInterval returns the cleanup interval.
	CleanupInterval() time.Duration

	// SetCleanupInterval sets the cleanup interval.
	SetCleanupInterval(value time.Duration) Options

	// CloseCheckInterval returns the close check interval.
	CloseCheckInterval() time.Duration

	// SetCloseCheckInterval sets the close check interval.
	SetCloseCheckInterval(value time.Duration) Options

	// InstrumentOptions returns the instrument options.
	InstrumentOptions() instrument.Options

	// SetInstrumentOptions sets the instrument options.
	SetInstrumentOptions(value instrument.Options) Options
}
