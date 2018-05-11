// +build integration

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

package integration

import (
	"testing"

	"github.com/m3db/m3msg/topic"
	"github.com/m3db/m3x/test"

	"github.com/fortytw2/leaktest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	maxProducers = 2
	maxRF        = 3
)

func TestSharedConsumer(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: 3},
		})

		s.Run(t, ctrl)
	}
}

func TestReplicatedConsumer(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: 3},
		})

		s.Run(t, ctrl)
	}
}

func TestSharedAndReplicatedConsumers(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		for j := 1; j <= maxRF; j++ {
			s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
				consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: j},
				consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: j},
			})

			s.Run(t, ctrl)
		}
	}
}

func TestSharedConsumerWithDeadInstance(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: 3},
		})

		pct := 10
		s.ScheduleOperations(
			pct,
			func() { s.KillInstance(t, 0) },
		)
		s.Run(t, ctrl)
		testConsumers := s.consumerServices[0].testConsumers
		require.True(t, testConsumers[len(testConsumers)-1].consumed <= s.TotalMessages()*pct/100)
	}
}

func TestSharedConsumerWithDeadConnection(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.KillConnection(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.KillConnection(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestReplicatedConsumerWithDeadConnection(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.KillConnection(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.KillConnection(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestSharedAndReplicatedConsumerWithDeadConnection(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		for j := 1; j <= maxRF; j++ {
			s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
				consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: j},
				consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: j},
			})

			s.ScheduleOperations(
				10,
				func() { s.KillConnection(t, 0) },
			)
			s.ScheduleOperations(
				20,
				func() { s.KillConnection(t, 1) },
			)
			s.ScheduleOperations(
				30,
				func() { s.KillConnection(t, 0) },
			)
			s.ScheduleOperations(
				40,
				func() { s.KillConnection(t, 1) },
			)
			s.Run(t, ctrl)
		}
	}
}

func TestSharedConsumerAddInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.AddInstance(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.AddInstance(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestReplicatedConsumerAddInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.AddInstance(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.AddInstance(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestSharedAndReplicatedConsumerAddInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		for j := 1; j <= maxRF; j++ {
			s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
				consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: j},
				consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: j},
			})

			s.ScheduleOperations(
				10,
				func() { s.AddInstance(t, 0) },
			)
			s.ScheduleOperations(
				20,
				func() { s.AddInstance(t, 1) },
			)
			s.ScheduleOperations(
				30,
				func() { s.AddInstance(t, 0) },
			)
			s.ScheduleOperations(
				40,
				func() { s.AddInstance(t, 1) },
			)
			s.Run(t, ctrl)
		}
	}
}

func TestSharedConsumerRemoveInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.RemoveInstance(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.RemoveInstance(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestReplicatedConsumerRemoveInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.RemoveInstance(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.RemoveInstance(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestSharedAndReplicatedConsumerRemoveInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		for j := 1; j <= maxRF; j++ {
			s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
				consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: j},
				consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: j},
			})

			s.ScheduleOperations(
				10,
				func() { s.RemoveInstance(t, 0) },
			)
			s.ScheduleOperations(
				20,
				func() { s.RemoveInstance(t, 1) },
			)
			s.ScheduleOperations(
				30,
				func() { s.RemoveInstance(t, 0) },
			)
			s.ScheduleOperations(
				40,
				func() { s.RemoveInstance(t, 1) },
			)
			s.Run(t, ctrl)
		}
	}
}

func TestSharedConsumerReplaceInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.ReplaceInstance(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.ReplaceInstance(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestReplicatedConsumerReplaceInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
			consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: 3},
		})

		s.ScheduleOperations(
			10,
			func() { s.ReplaceInstance(t, 0) },
		)
		s.ScheduleOperations(
			20,
			func() { s.ReplaceInstance(t, 0) },
		)
		s.Run(t, ctrl)
	}
}

func TestSharedAndReplicatedConsumerReplaceInstances(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	defer leaktest.Check(t)()

	ctrl := gomock.NewController(test.Reporter{t})
	defer ctrl.Finish()

	for i := 1; i <= maxProducers; i++ {
		for j := 1; j <= maxRF; j++ {
			s := newTestSetup(t, ctrl, i, []consumerServiceConfig{
				consumerServiceConfig{ct: topic.Shared, instances: 5, replicas: j},
				consumerServiceConfig{ct: topic.Replicated, instances: 5, replicas: j},
			})

			s.ScheduleOperations(
				10,
				func() { s.ReplaceInstance(t, 0) },
			)
			s.ScheduleOperations(
				20,
				func() { s.ReplaceInstance(t, 1) },
			)
			s.ScheduleOperations(
				30,
				func() { s.ReplaceInstance(t, 0) },
			)
			s.ScheduleOperations(
				40,
				func() { s.ReplaceInstance(t, 1) },
			)
			s.Run(t, ctrl)
		}
	}
}
