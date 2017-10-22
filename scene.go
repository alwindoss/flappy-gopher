package main

import (
	"context"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	time  int
	bg    *sdl.Texture
	birds []*sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("unable to load background texture: %v", err)
	}
	var birds []*sdl.Texture
	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("res/imgs/bird_frame_%d.png", i)
		bird, err := img.LoadTexture(r, path)
		if err != nil {
			return nil, fmt.Errorf("unable to load bird texture: %v", err)
		}
		birds = append(birds, bird)
	}

	return &scene{bg: bg, birds: birds}, err
}

func (s *scene) run(ctx context.Context, r *sdl.Renderer) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		for range time.Tick(100 * time.Millisecond) {
			select {
			case <-ctx.Done():
				return
			default:
				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()
	return errc
}

func (s *scene) paint(r *sdl.Renderer) error {
	s.time++
	r.Clear()
	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("unable to paint background: %v", err)
	}

	rect := &sdl.Rect{H: 43, W: 50, X: 10, Y: 300 - 43/2}
	i := s.time % len(s.birds)
	if err := r.Copy(s.birds[i], nil, rect); err != nil {
		return fmt.Errorf("unable to paint bird: %v", err)
	}
	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
