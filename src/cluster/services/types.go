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

package services

import (
	"time"

	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3x/instrument"
	xwatch "github.com/m3db/m3x/watch"
)

// Services provides access to the service topology
type Services interface {
	// Advertise advertises the availability of an instance of a service
	Advertise(ad Advertisement) error

	// Unadvertise indicates a given instance is no longer available
	Unadvertise(service ServiceID, id string) error

	// Query returns the topology for a given service
	Query(service ServiceID, opts QueryOptions) (Service, error)

	// Watch returns a watch on metadata and a list of available instances for a given service
	Watch(service ServiceID, opts QueryOptions) (xwatch.Watch, error)

	// Metadata returns the metadata for a given service
	Metadata(sid ServiceID) (Metadata, error)

	// SetMetadata sets the metadata for a given service
	SetMetadata(sid ServiceID, m Metadata) error

	// PlacementService returns a client of Placement Service
	PlacementService(service ServiceID, popts PlacementOptions) (PlacementService, error)

	// HeartbeatService returns a heartbeat store for the given service.
	HeartbeatService(service ServiceID) (HeartbeatService, error)
}

// Service describes the metadata and instances of a service
type Service interface {
	// Instance returns the service instance with the instance id
	Instance(instanceID string) (ServiceInstance, error)

	// Instances returns the service instances
	Instances() []ServiceInstance

	// SetInstances sets the service instances
	SetInstances(insts []ServiceInstance) Service

	// Replication returns the service replication description or nil if none
	Replication() ServiceReplication

	// SetReplication sets the service replication description or nil if none
	SetReplication(r ServiceReplication) Service

	// Sharding returns the service sharding description or nil if none
	Sharding() ServiceSharding

	// SetSharding sets the service sharding description or nil if none
	SetSharding(s ServiceSharding) Service
}

// ServiceReplication describes the replication of a service
type ServiceReplication interface {
	// Replicas is the count of replicas
	Replicas() int

	// SetReplicas sets the count of replicas
	SetReplicas(r int) ServiceReplication
}

// ServiceSharding describes the sharding of a service
type ServiceSharding interface {
	// NumShards is the number of shards to use for sharding
	NumShards() int

	// SetNumShards sets the number of shards to use for sharding
	SetNumShards(n int) ServiceSharding

	// IsSharded() returns whether this service is sharded
	IsSharded() bool

	// SetIsSharded sets IsSharded
	SetIsSharded(s bool) ServiceSharding
}

// ServiceInstance is a single instance of a service
type ServiceInstance interface {
	ServiceID() ServiceID                           // the service implemented by the instance
	SetServiceID(service ServiceID) ServiceInstance // sets the service implemented by the instance
	InstanceID() string                             // ID of the instance
	SetInstanceID(id string) ServiceInstance        // sets the ID of the instance
	Endpoint() string                               // Endpoint address for contacting the instance
	SetEndpoint(e string) ServiceInstance           // sets the endpoint address for the instance
	Shards() shard.Shards                           // Shards owned by the instance
	SetShards(s shard.Shards) ServiceInstance       // sets the Shards assigned to the instance
}

// Advertisement advertises the availability of a given instance of a service
type Advertisement interface {
	ServiceID() ServiceID                         // the service being advertised
	SetServiceID(service ServiceID) Advertisement // sets the service being advertised
	Health() func() error                         // optional health function, return an error to indicate unhealthy
	SetHealth(health func() error) Advertisement  // sets the health function for the advertised instance

	// Returns the placement instance associated with this advertisement, which
	// contains the ID of the instance advertising and all other relevant fields.
	PlacementInstance() PlacementInstance

	// Sets the PlacementInstance that is advertising.
	SetPlacementInstance(p PlacementInstance) Advertisement
}

// ServiceID contains the fields required to id a service
type ServiceID interface {
	String() string                      // String returns a description of the ServiceID
	Name() string                        // the service name of the ServiceID
	SetName(s string) ServiceID          // set the service name of the ServiceID
	Environment() string                 // the environemnt of the ServiceID
	SetEnvironment(env string) ServiceID // sets the environemnt of the ServiceID
	Zone() string                        // the zone of the ServiceID
	SetZone(zone string) ServiceID       // sets the zone of the ServiceID
}

// QueryOptions are options to service discovery queries
type QueryOptions interface {
	IncludeUnhealthy() bool                  // if true, will return unhealthy instances
	SetIncludeUnhealthy(h bool) QueryOptions // sets whether to include unhealthy instances
}

// Metadata contains the metadata for a service
type Metadata interface {
	// String returns a description of the metadata
	String() string

	// Port returns the port to be used to contact the service
	Port() uint32

	// SetPort sets the port
	SetPort(p uint32) Metadata

	// LivenessInterval is the ttl interval for an instance to be considered as healthy
	LivenessInterval() time.Duration

	// SetLivenessInterval sets the LivenessInterval
	SetLivenessInterval(l time.Duration) Metadata

	// HeartbeatInterval is the interval for heatbeats
	HeartbeatInterval() time.Duration

	// SetHeartbeatInterval sets the HeartbeatInterval
	SetHeartbeatInterval(h time.Duration) Metadata
}

// PlacementService handles the placement related operations for registered services
// all write or update operations will persist the generated placement before returning success
type PlacementService interface {
	// BuildInitialPlacement initialize a placement
	BuildInitialPlacement(instances []PlacementInstance, numShards int, rf int) (ServicePlacement, error)

	// AddReplica up the replica factor by 1 in the placement
	AddReplica() (ServicePlacement, error)

	// AddInstance picks an instance from the candidate list to the placement
	AddInstance(candidates []PlacementInstance) (ServicePlacement, error)

	// RemoveInstance removes an instance from the placement
	RemoveInstance(leavingInstanceID string) (ServicePlacement, error)

	// ReplaceInstance picks instances from the candidate list to replace an instance in current placement
	ReplaceInstance(leavingInstanceID string, candidates []PlacementInstance) (ServicePlacement, error)

	// MarkShardAvailable marks the state of a shard as available
	MarkShardAvailable(instanceID string, shardID uint32) error

	// MarkInstanceAvailable marks all the shards on a given instance as available
	MarkInstanceAvailable(instanceID string) error

	// Placement returns the persisted placement with version
	Placement() (ServicePlacement, int, error)

	// SetPlacement persists the placement
	SetPlacement(p ServicePlacement) error

	// Delete deletes the placement
	Delete() error
}

// PlacementOptions is the interface for placement options
type PlacementOptions interface {
	// LooseRackCheck enables the placement to loose the rack check
	// during instance replacement to achieve full ownership transfer
	LooseRackCheck() bool
	SetLooseRackCheck(looseRackCheck bool) PlacementOptions

	// AllowPartialReplace allows shards from the leaving instance to be
	// placed on instances other than the new instances in a replace operation
	AllowPartialReplace() bool
	SetAllowPartialReplace(allowPartialReplace bool) PlacementOptions

	// IsSharded describes whether a placement needs to be sharded
	// when set to false, no specific shards will be assigned to any instance
	IsSharded() bool
	SetIsSharded(sharded bool) PlacementOptions

	// Dryrun will try to perform the placement operation but will not persist the final result
	Dryrun() bool
	SetDryrun(d bool) PlacementOptions

	// InstrumentOptions is the options for instrument
	InstrumentOptions() instrument.Options
	SetInstrumentOptions(iopts instrument.Options) PlacementOptions

	// ValidZone returns the zone that added instances must be in in order
	// to be added to a placement.
	ValidZone() string

	// SetValidZone sets the zone that added instances must be in in order to
	// be added to a placement. By default the valid zone will be the zone of
	// instances already in a placement, however if a placement is empty then
	// it is necessary to specify the valid zone when adding the first
	// instance.
	SetValidZone(z string) PlacementOptions
}

// ServicePlacement describes how instances are placed in a service
type ServicePlacement interface {
	// Instances returns all Instances in the placement
	Instances() []PlacementInstance

	// SetInstances sets the Instances
	SetInstances(instances []PlacementInstance) ServicePlacement

	// NumInstances returns the number of instances in the placement
	NumInstances() int

	// Instance returns the Instance for the requested id
	Instance(id string) (PlacementInstance, bool)

	// ReplicaFactor returns the replica factor in the placement
	ReplicaFactor() int

	// SetReplicaFactor sets the ReplicaFactor
	SetReplicaFactor(rf int) ServicePlacement

	// Shards returns all the unique shard ids for a replica
	Shards() []uint32

	// SetShards sets the unique shard ids for a replica
	SetShards(s []uint32) ServicePlacement

	// ShardsLen returns the number of shards in a replica
	NumShards() int

	// IsSharded() returns whether this placement is sharded
	IsSharded() bool

	// SetIsSharded() sets IsSharded
	SetIsSharded(v bool) ServicePlacement

	// String returns a description of the placement
	String() string

	// GetVersion() returns the version of the placement retreived from the
	// backing MVCC store.
	GetVersion() int

	// SetVersion() sets the version of the placement object. Since version
	// is determined by the backing MVCC store, calling this method has no
	// effect in terms of the updated ServicePlacement that is written back
	// to the MVCC store.
	SetVersion(v int) ServicePlacement
}

// PlacementInstance represents an instance in a service placement
type PlacementInstance interface {
	String() string                             // String is for debugging
	ID() string                                 // ID is the id of the instance
	SetID(id string) PlacementInstance          // SetID sets the id of the instance
	Rack() string                               // Rack is the rack of the instance
	SetRack(r string) PlacementInstance         // SetRack sets the rack of the instance
	Zone() string                               // Zone is the zone of the instance
	SetZone(z string) PlacementInstance         // SetZone sets the zone of the instance
	Weight() uint32                             // Weight is the weight of the instance
	SetWeight(w uint32) PlacementInstance       // SetWeight sets the weight of the instance
	Endpoint() string                           // Endpoint is the endpoint of the instance
	SetEndpoint(ip string) PlacementInstance    // SetEndpoint sets the endpoint of the instance
	Shards() shard.Shards                       // Shards returns the shards owned by the instance
	SetShards(s shard.Shards) PlacementInstance // SetShards sets the shards owned by the instance
}

// HeartbeatService manages heartbeating instances
type HeartbeatService interface {
	// Heartbeat sends heartbeat for a service instance with a ttl
	Heartbeat(instance PlacementInstance, ttl time.Duration) error

	// Get gets healthy instances for a service
	Get() ([]string, error)

	// GetInstances returns a deserialized list of healthy PlacementInstances.
	GetInstances() ([]PlacementInstance, error)

	// Delete deletes the heartbeat for a service instance
	Delete(instance string) error

	// Watch watches the heartbeats for a service
	Watch() (xwatch.Watch, error)
}
