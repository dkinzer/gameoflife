package cell

import(
  "fmt"
  "reflect"
  "testing"
  p "github.com/dkinzer/gameoflife/point"
)

func TestFactorySinglePoint(t *testing.T) {
  expected := map[p.Point]Cell{
    p.Point{1, 2}:Cell{
      Point: p.Point{1, 2},
      Status: ALIVE,
      LiveNeighbors: []p.Point{},
    },
  }
  actual := CellFactory([]p.Point{
    p.Point{1, 2},
  })

  if !reflect.DeepEqual(actual, expected) {
    t.Error("CellFactory(Single Point slice).")
  }
}

func TestFactoryTwoSeperatePoints(t *testing.T) {
  expected := map[p.Point]Cell{
    p.Point{1, 2}: Cell{
      Point: p.Point{1, 2},
      Status: ALIVE,
      LiveNeighbors: []p.Point{},
    },
    p.Point{4, 5}: Cell{
      Point: p.Point{4, 5},
      Status: ALIVE,
      LiveNeighbors: []p.Point{},
    },
  }
  actual := CellFactory([]p.Point{
    p.Point{1, 2},
    p.Point{4, 5},
  })

  if !reflect.DeepEqual(actual, expected) {
    t.Error("CellFactory(Two seperate Point slice).")
  }
}

func TestFactoryTwoJoinedPoints(t *testing.T) {
  expected := map[p.Point]Cell{
    p.Point{1, 2}: Cell{
      Point: p.Point{1, 2},
      Status: ALIVE,
      LiveNeighbors: []p.Point{
        p.Point{2, 3},
      },
    },
    p.Point{2, 3}: Cell{
      Point: p.Point{2, 3},
      Status: ALIVE,
      LiveNeighbors: []p.Point{
        p.Point{1, 2},
      },
    },
  }
  actual := CellFactory([]p.Point{
    p.Point{1, 2},
    p.Point{2, 3},
  })

  if !reflect.DeepEqual(actual, expected) {
    fmt.Println("expected: ", expected)
    fmt.Println("actual: ", actual)
    t.Error("CellFactory(Two joined Point slice).")
  }
}
