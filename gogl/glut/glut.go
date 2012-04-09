package glut
/*
#cgo pkg-config: gl
#cgo LDFLAGS: -lglut
#include <GL/glut.h>
#include <GL/gl.h>

static float rotation = 0.0f;

void disp(void) {
	// do  a clearscreen
	glClear(GL_COLOR_BUFFER_BIT);

	glPushMatrix();

	rotation += 5.0;
	glutPostRedisplay();

	glRotatef(rotation, 1.0, 0.0, 0.0);
	glRotatef(rotation/2.0, 0.0, -1.0, 0.0);
	glRotatef(rotation/3.0, 0.0, 0.0, 1.0);
	glutWireTeapot(0.5);
	glPopMatrix();

	glutSwapBuffers();
}

void Stuff() {
	glutDisplayFunc(disp);
	glClearColor(0.0,0.0,0.0,0.0);
	glutMainLoop();
}
*/
import "C"
import "unsafe"

const (
	GLUT_RGBA = C.GLUT_RGBA
	GLUT_DOUBLE = C.GLUT_DOUBLE
)

func Init() {
	var size C.int = 0
	C.glutInit(&size, nil)
}

func InitDisplayMode(mode int) {
	C.glutInitDisplayMode(C.uint(mode))
}

func InitWindowSize(width int, height int) {
	C.glutInitWindowSize(C.int(width), C.int(height))
}

func CreateWindow(name string) {
	cstr := C.CString(name)
	C.glutCreateWindow(cstr)
	C.free(unsafe.Pointer(cstr))
}

func Stuff() {
	C.Stuff();
}

