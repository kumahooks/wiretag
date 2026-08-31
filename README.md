Introduction
============

[![Go Reference](https://pkg.go.dev/badge/github.com/kumahooks/wiretag.svg)](https://pkg.go.dev/github.com/kumahooks/wiretag)

The wiretag package is an idiomatic Go bindings for [TagLib](https://taglib.org/). wiretag wraps the C bindings taglib itself ships (`tag_c.h`) with Go-native error handling and explicit memory ownership.

This release was built and tested against taglib 2.3.1.

Installation
============

Dependencies
------------
You need the taglib C bindings installed on the system before building or running anything that imports wiretag:

- Debian / Ubuntu: `apt install libtagc0-dev` (or `libtag-c-dev` on newer releases)
- macOS: `brew install taglib`

Install the library itself:

```sh
go get github.com/kumahooks/wiretag
```

Usage
============

Example
------------
Open a file, read its tags and properties, and close it:

```go
package main

import (
	"fmt"
	"log"

	"github.com/kumahooks/wiretag"
)

func main() {
	file, err := wiretag.Open("song.flac")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	title, err := file.Title()
	if err != nil {
		log.Fatal(err)
	}

	properties, err := file.Properties()
	if err != nil {
		log.Fatal(err)
	}

	channels, err := file.AudioChannels()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(title)      // "Ketsarku Mozgalom"
	fmt.Println(properties) // map[ALBUM:[... ] ARTIST:[...] ...]
	fmt.Println(channels)   // 2
}
```

Under Development!
------------
wiretag 0.1 is read-only. Writing tags, saving files, memory-backed iostreams, and complex properties (such as embedded pictures) are not implemented yet and are planned for upcoming releases.

License
============

BSD 3-Clause. See [LICENSE](LICENSE).

