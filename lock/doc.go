// Package lock defines interfaces for a priority-based lock queue in RoadRunner.
// The Queue interface manages lock items ordered by priority, supporting Insert,
// Remove, and ExtractMin operations. Each Item carries an ID, GroupID, and
// Priority used for ordering within the queue.
package lock
