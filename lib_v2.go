package cgomultiversion

// #cgo LDFLAGS: -l:libcgomultiversion.so.2
// #include <cgomultiversion_v2/lib.h>
import "C"

func AddV2(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
