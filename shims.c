// shims.c defines C helper functions for some specific taglib functionalities.
#include <taglib/tag_c.h>

const char* wiretag_strarray_at(char** array, int index)
{
	return array[index];
}

