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
    t.Error("CellFactory(Single Point slice).")
  }
}

func TestFactoryTwoSeperatePoints(t *testing.T) {
  m := make(map[p.Point]Cell, 2)
  m[p.Point{1, 2}] = Cell{
    Point: p.Point{1, 2},
    Status: ALIVE,
    LiveNeighbors: []p.Point{},
  }

  m[p.Point{4, 5}] = Cell{
    Point: p.Point{4, 5},
    Status: ALIVE,
    LiveNeighbors: []p.Point{},
  }
  expected := m
  actual := CellFactory([]p.Point{
    p.Point{1, 2},
    p.Point{4, 5},
  })

  if !reflect.DeepEqual(actual, expected) {
    t.Error("CellFactory(Two seperate Point slice).")
  }
}

func TestFactoryTwoJoinedPoints(t *testing.T) {
  m := make(map[p.Point]Cell, 2)
  m[p.Point{1, 2}] = Cell{
    Point: p.Point{1, 2},
    Status: ALIVE,
    LiveNeighbors: []p.Point{
      p.Point{2, 3},
    },
  }
  m[p.Point{2, 3}] = Cell{
    Point: p.Point{2, 3},
    Status: ALIVE,
    LiveNeighbors: []p.Point{
      p.Point{1, 2},
    },
  }

  expected := m
  actual := CellFactory([]p.Point{
    p.Point{1, 2},
    p.Point{2, 3},
  })

  if !reflect.DeepEqual(actual, expected) {
    t.Error("CellFactory(Two joined Point slice).")
  }
}
