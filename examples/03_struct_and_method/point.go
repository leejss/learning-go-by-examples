package main

import "math"

type Point struct {
	X, Y float64
}

func (p Point) DistanceTo(q Point) float64 {
	dx := p.X - q.X
	dy := p.Y - q.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p *Point) Move(dx, dy float64) *Point {
	p.X += dx
	p.Y += dy
	return p
}
