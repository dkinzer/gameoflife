package cell

import(
  "reflect"
  "testing"
  p "github.com/dkinzer/gameoflife/point"
)

func TestFactorySinglePoint(t *testing.T) {
  m := make(map[p.Point]Cell, 1)
  m[p.Point{1, 2}] = Cell{
    Point: p.Point{1, 2},
    Status: ALIVE,
    LiveNeighbors: []p.Point{},
  }

  expected := m
  actual := CellFactory([]p.Point{
    p.Point{1, 2},
  })

  if !reflect.DeepEqual(actual, expected) {
    t.Error("CellFactory(Single Point Slice).")
  }
}

