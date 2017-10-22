package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/veandco/go-sdl2/sdl"
)

func run() error {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()
	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not initialize ttf: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create Window and renderer: %v", err)
	}
	defer w.Destroy()

	if err := drawTile(r); err != nil {
		return fmt.Errorf("could not draw title %v", err)
	}
	time.Sleep(5 * time.Second)

	if err := drawBackground(r); err != nil {
		return fmt.Errorf("unable to draw background %v", err)
	}
	time.Sleep(5 * time.Second)

	return err
}

func drawBackground(r *sdl.Renderer) error {
	r.Clear()
	// img.Init(img.INIT_PNG)
	t, err := img.LoadTexture(r, "res/imgs/background.png")
	if err != nil {
		return fmt.Errorf("unable to load texture: %v", err)
	}
	defer t.Destroy()
	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("unable to copy texture into the rendered: %v", err)
	}
	r.Present()
	return err
}

func drawTile(r *sdl.Renderer) error {
	if err := r.Clear(); err != nil {
		return fmt.Errorf("unable to clear %v", err)
	}
	fnt, err := ttf.OpenFont("res/fonts/Flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("unable to open font: %v", err)
	}
	defer fnt.Close()

	c := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	s, err := fnt.RenderUTF8_Solid("Flappy Gopher", c)
	if err != nil {
		return fmt.Errorf("unable to Render Solid: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("unable to create Texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("unable to Copy %v", err)
	}
	r.Present()
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}
