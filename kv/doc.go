// Package kv defines interfaces for key-value storage in RoadRunner.
// The Storage interface provides operations such as Get, Set, Delete, and TTL
// management across pluggable backends. Item represents a single entry with a
// key, value, and optional timeout. The Constructor interface allows drivers to
// be instantiated from configuration.
package kv
