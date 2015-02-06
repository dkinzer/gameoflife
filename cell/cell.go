package cell

import (
  p "github.com/dkinzer/gameoflife/point"
)

type status byte

const (
  ALIVE status = iota
  DEAD
)

type Cell struct {
  p.Point
  Status status
  LiveNeighbors []p.Point
}

func (c Cell) nextGeneration() Cell {
  n := len(c.LiveNeighbors)
  switch {
  case n < 2:
    c.Status = DEAD

  case n == 3:
    c.Status = ALIVE

  case n > 3:
    c.Status = DEAD
  }
  return c
}

