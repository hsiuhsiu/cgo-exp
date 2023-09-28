package mycgolib

/*
#cgo CFLAGS: -I${SRCDIR}/../clib
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/../clib
#cgo LDFLAGS: -L${SRCDIR}/../clib
#cgo LDFLAGS: -lpoint

#include <point.h>
*/
import "C"

func Sum(a, b int) int {
	return int(C.sum(C.int(a), C.int(b)))
}
