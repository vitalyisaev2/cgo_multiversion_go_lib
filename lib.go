package cgomultiversion

// #cgo LDFLAGS: -l:libcgomultiversion.so.1
// #include <cgomultiversion_v1/lib.h>
import "C"

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
