package main

import "glut"

func main() {
	glut.Init();
	glut.InitDisplayMode(glut.GLUT_RGBA | glut.GLUT_DOUBLE)
	glut.InitWindowSize(1024, 480)
	glut.CreateWindow("Test")

	glut.Stuff()
}

