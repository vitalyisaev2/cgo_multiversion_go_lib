package cgomultiversion

// #cgo LDFLAGS: -l:libcgomultiversion.so.1
// #include <cgomultiversion_v1/lib.h>
import "C"

func AddV1(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
