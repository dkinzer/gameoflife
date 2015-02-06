package cell

import(
  p "github.com/dkinzer/gameoflife/point"
)

func CellFactory(points []p.Point) map[p.Point]Cell {
  cells := make(map[p.Point]Cell, len(points))

  for _, point := range(points) {
    cells[point] = Cell{
      Point: point,
      Status: ALIVE,
      LiveNeighbors: []p.Point{},
    }
  }
  return cells
}

