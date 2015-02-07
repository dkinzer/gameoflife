package cell

import(
  p "github.com/dkinzer/gameoflife/point"
)

func CellFactory(points []p.Point) map[p.Point]Cell {
  cells := make(map[p.Point]Cell, len(points))

  for _, point := range points {
    cells[point] = Cell{
      Point: point,
      Status: ALIVE,

      LiveNeighbors: (func (point p.Point) []p.Point {
        neighbors := make([]p.Point, 0)
        for _, neighbor := range point.Neighbors() {
          for _, node := range points {
            if neighbor == node {
              neighbors= append(neighbors, neighbor)
            }
          }
        }
        return neighbors
      })(point),
    }
  }
  return cells
}

