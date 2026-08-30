package wiretag

// #cgo pkg-config: taglib
// #cgo LDFLAGS: -ltag_c
// #include <taglib/tag_c.h>
//
// extern void wiretag_init(void);
import "C"

func init() {
	C.wiretag_init()
}
