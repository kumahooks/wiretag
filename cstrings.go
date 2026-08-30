package wiretag

// #include <taglib/tag_c.h>
//
// extern const char* wiretag_strarray_at(char** array, int index);
import "C"

// cStringSlice copies a NULL-terminated char** into a Go string slice.
func cStringSlice(arr **C.char) []string {
	values := []string{}
	for i := 0; ; i++ {
		value := C.wiretag_strarray_at(arr, C.int(i))
		if value == nil {
			break
		}

		values = append(values, C.GoString(value))
	}

	return values
}
