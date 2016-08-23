package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

var (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatalln("Could not initialize GLFW:", err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Resizable, glfw.False)

	// Create window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Shapes muh-fuckah", nil, nil)
	if err != nil {
		log.Fatalln("Could not create window:", err)
	}

	// Setup OpenGL context
	window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		log.Fatalln("Could not initialize OpenGL:", err)
	}

	window.SetKeyCallback(keyCB)
	// Event loop
	glfw.SwapInterval(1)
	for !window.ShouldClose() {
		width, height := window.GetFramebufferSize()
		gl.Viewport(0, 0, int32(width), int32(height))

		glfw.PollEvents()
		window.SwapBuffers()
	}

}

func keyCB(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mode glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Release {
		window.SetShouldClose(true)
	}

	if key == glfw.KeyB && action == glfw.Press {
		log.Println("Bedtime")
	}
}
