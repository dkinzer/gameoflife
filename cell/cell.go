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

func (c Cell) nextGenerationStatus() status {
  n := len(c.LiveNeighbors)
  switch {
  case n < 2:
    return DEAD

  case n == 3:
    return ALIVE

  case n > 3:
    return DEAD
  }
  return c.Status
}

