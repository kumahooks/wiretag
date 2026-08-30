// Package wiretag exposes Go bindings for popular taglib C++ library.
package wiretag

// #cgo pkg-config: taglib
// #include <taglib/tag_c.h>
import "C"

func Version() string {
	return "2.3.1"
}
