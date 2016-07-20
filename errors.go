package orb

import "errors"

var (
	// ErrUnsupportedDataType is returned by Scan methods when asked to scan
	// non []byte data from the database. This should never happen
	// if the driver is acting appropriately.
	ErrUnsupportedDataType = errors.New("orb: scan value must be []byte")

	// ErrNotWKB is returned when unmarshalling WKB and the data is not valid.
	ErrNotWKB = errors.New("orb: invalid WKB data")

	// ErrIncorrectGeometry is returned when unmarshalling WKB data into the wrong type.
	// For example, unmarshaling linestring data into a point.
	ErrIncorrectGeometry = errors.New("orb: incorrect geometry")
)

var (
	// ErrPointOutsideOfBounds is returned when trying to add a point
	// to a quad tree and the point is outside the bounds used to create the tree.
	ErrPointOutsideOfBounds = errors.New("quadtree: point outside of bounds")
)
