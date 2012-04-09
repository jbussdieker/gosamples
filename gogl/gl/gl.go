package gl
/*
#cgo pkg-config: gl
#include <GL/gl.h>
*/
import "C"

const (
	GL_COLOR_BUFFER_BIT = C.GL_COLOR_BUFFER_BIT
)

func Clear(mode int) {
	C.glClear(C.GLbitfield(mode))
}

func ClearColor(r float64, g float64, b float64, a float64) {
	C.glClearColor(r,g,b,a)
}
