package cell

import(
  "testing"
  p "github.com/dkinzer/gameoflife/point"
)

func TestFactorySinglePoint(t *testing.T) {
  m := make(map[p.Point]Cell)
  m[p.Point{1, 2}] = Cell{
    Point: p.Point{1, 2},
    Status: ALIVE,
    LiveNeighbors: []p.Point{},
  }

  expected := m
  actual := CellFactory([]p.Point{
    p.Point{1, 2},
  })

  if !compareMapPointCell(actual, expected) {
    t.Error("CellFactory(Single Point Slice).")
  }
}

func compareMapPointCell(m1, m2 map[p.Point]Cell) bool {
  if len(m1) != len(m2) {
    return false
  }
  for point, _ := range(m1) {
    if !compareCell(m1[point], m2[point]) {
      return false
    }
  }
  return true
}

func compareCell(c1, c2 Cell) bool {
  return c1.Point == c2.Point &&
  c1.Status == c2.Status &&
  comparePointSlice(c1.LiveNeighbors, c2.LiveNeighbors)
}

func comparePointSlice(s1, s2 []p.Point) bool {
  if len(s1) != len(s2) {
    return false
  }
  for i, _ := range(s2) {
    if s1[i] != s2[i] {
      return false
    }
  }
  return true
}

