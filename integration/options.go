// Copyright (c) 2016 Uber Technologies, Inc.
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

package integration

import (
	"time"
)

const (
	// defaultServerStateChangeTimeout is the default time we wait for a server to change its state.
	defaultServerStateChangeTimeout = time.Minute

	// defaultClusterConnectionTimeout is the default time we wait for cluster connections to be established.
	defaultClusterConnectionTimeout = 2 * time.Second

	// defaultReadRequestTimeout is the default read request timeout.
	defaultReadRequestTimeout = 2 * time.Second

	// defaultWriteRequestTimeout is the default write request timeout.
	defaultWriteRequestTimeout = 2 * time.Second

	// defaultTruncateRequestTimeout is the default truncate request timeout.
	defaultTruncateRequestTimeout = 2 * time.Second

	// defaultWorkerPoolSize is the default number of workers in the worker pool.
	defaultWorkerPoolSize = 10

	// defaultUseTChannelClientForReading determines whether we use the tchannel client for reading by default.
	defaultUseTChannelClientForReading = true

	// defaultUseTChannelClientForWriting determines whether we use the tchannel client for writing by default.
	defaultUseTChannelClientForWriting = false

	// defaultUseTChannelClientForTruncation determines whether we use the tchannel client for truncation by default.
	defaultUseTChannelClientForTruncation = true
)

type testOptions interface {
	// SetServerStateChangeTimeout sets the server state change timeout.
	SetServerStateChangeTimeout(value time.Duration) testOptions

	// ServerStateChangeTimeout returns the server state change timeout.
	ServerStateChangeTimeout() time.Duration

	// SetClusterConnectionTimeout sets the cluster connection timeout.
	SetClusterConnectionTimeout(value time.Duration) testOptions

	// ClusterConnectionTimeout returns the cluster connection timeout.
	ClusterConnectionTimeout() time.Duration

	// SetReadRequestTimeout sets the read request timeout.
	SetReadRequestTimeout(value time.Duration) testOptions

	// ReadRequestTimeout returns the read request timeout.
	ReadRequestTimeout() time.Duration

	// SetWriteRequestTimeout sets the write request timeout.
	SetWriteRequestTimeout(value time.Duration) testOptions

	// WriteRequestTimeout returns the write request timeout.
	WriteRequestTimeout() time.Duration

	// SetTruncateRequestTimeout sets the truncate request timeout.
	SetTruncateRequestTimeout(value time.Duration) testOptions

	// TruncateRequestTimeout returns the truncate request timeout.
	TruncateRequestTimeout() time.Duration

	// SetWorkerPoolSize sets the number of workers in the worker pool.
	SetWorkerPoolSize(value int) testOptions

	// WorkerPoolSize returns the number of workers in the worker pool.
	WorkerPoolSize() int

	// SetUseTChannelClientForReading sets whether we use the tchannel client for reading.
	SetUseTChannelClientForReading(value bool) testOptions

	// UseTChannelClientForReading returns whether we use the tchannel client for reading.
	UseTChannelClientForReading() bool

	// SetUseTChannelClientForWriting sets whether we use the tchannel client for writing.
	SetUseTChannelClientForWriting(value bool) testOptions

	// UseTChannelClientForWriting returns whether we use the tchannel client for writing.
	UseTChannelClientForWriting() bool

	// SetUseTChannelClientForTruncation sets whether we use the tchannel client for truncation.
	SetUseTChannelClientForTruncation(value bool) testOptions

	// UseTChannelClientForTruncation returns whether we use the tchannel client for truncation.
	UseTChannelClientForTruncation() bool
}

type options struct {
	serverStateChangeTimeout       time.Duration
	clusterConnectionTimeout       time.Duration
	readRequestTimeout             time.Duration
	writeRequestTimeout            time.Duration
	truncateRequestTimeout         time.Duration
	workerPoolSize                 int
	useTChannelClientForReading    bool
	useTChannelClientForWriting    bool
	useTChannelClientForTruncation bool
}

func newTestOptions() testOptions {
	return &options{
		serverStateChangeTimeout:       defaultServerStateChangeTimeout,
		clusterConnectionTimeout:       defaultClusterConnectionTimeout,
		readRequestTimeout:             defaultReadRequestTimeout,
		writeRequestTimeout:            defaultWriteRequestTimeout,
		truncateRequestTimeout:         defaultTruncateRequestTimeout,
		workerPoolSize:                 defaultWorkerPoolSize,
		useTChannelClientForReading:    defaultUseTChannelClientForReading,
		useTChannelClientForWriting:    defaultUseTChannelClientForWriting,
		useTChannelClientForTruncation: defaultUseTChannelClientForTruncation,
	}
}

func (o *options) SetServerStateChangeTimeout(value time.Duration) testOptions {
	opts := *o
	opts.serverStateChangeTimeout = value
	return &opts
}

func (o *options) ServerStateChangeTimeout() time.Duration {
	return o.serverStateChangeTimeout
}

func (o *options) SetClusterConnectionTimeout(value time.Duration) testOptions {
	opts := *o
	opts.clusterConnectionTimeout = value
	return &opts
}

func (o *options) ClusterConnectionTimeout() time.Duration {
	return o.clusterConnectionTimeout
}

func (o *options) SetReadRequestTimeout(value time.Duration) testOptions {
	opts := *o
	opts.readRequestTimeout = value
	return &opts
}

func (o *options) ReadRequestTimeout() time.Duration {
	return o.readRequestTimeout
}

func (o *options) SetWriteRequestTimeout(value time.Duration) testOptions {
	opts := *o
	opts.writeRequestTimeout = value
	return &opts
}

func (o *options) WriteRequestTimeout() time.Duration {
	return o.writeRequestTimeout
}

func (o *options) SetTruncateRequestTimeout(value time.Duration) testOptions {
	opts := *o
	opts.truncateRequestTimeout = value
	return &opts
}

func (o *options) TruncateRequestTimeout() time.Duration {
	return o.truncateRequestTimeout
}

func (o *options) SetWorkerPoolSize(value int) testOptions {
	opts := *o
	opts.workerPoolSize = value
	return &opts
}

func (o *options) WorkerPoolSize() int {
	return o.workerPoolSize
}

func (o *options) SetUseTChannelClientForReading(value bool) testOptions {
	opts := *o
	opts.useTChannelClientForReading = value
	return &opts
}

func (o *options) UseTChannelClientForReading() bool {
	return o.useTChannelClientForReading
}

func (o *options) SetUseTChannelClientForWriting(value bool) testOptions {
	opts := *o
	opts.useTChannelClientForWriting = value
	return &opts
}

func (o *options) UseTChannelClientForWriting() bool {
	return o.useTChannelClientForWriting
}

func (o *options) SetUseTChannelClientForTruncation(value bool) testOptions {
	opts := *o
	opts.useTChannelClientForTruncation = value
	return &opts
}

func (o *options) UseTChannelClientForTruncation() bool {
	return o.useTChannelClientForTruncation
}