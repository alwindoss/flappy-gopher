package main

import (
	"context"
	"fmt"
	"os"
	"time"

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
	time.Sleep(1 * time.Second)

	s, err := newScene(r)
	if err != nil {
		return fmt.Errorf("unable to create a new scene: %v", err)
	}
	defer s.destroy()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	select {
	case err := <-s.run(ctx, r):
		return err
	case <-time.After(5 * time.Second):
		return nil
	}

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
