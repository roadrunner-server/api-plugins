package lock

// Item interface represents the base meta-information which any priority queue message must have
type Item interface {
	// ID returns a unique identifier for the item
	ID() string
	// GroupID returns the group associated with the item, used to remove all items with the same groupID
	GroupID() string
	// Priority returns the priority level used to sort the item
	Priority() int64
}

// Queue represents Lock plugin queue with it's elements types inside
type Queue interface {
	// Remove removes element with provided ID (if exists) and returns that elements
	Remove(id string) []Item
	// Insert adds an item to the queue
	Insert(item Item)
	// ExtractMin returns the item with the highest priority (less value is the highest priority)
	ExtractMin() Item
	// Len returns the number of items in the queue
	Len() uint64
}
