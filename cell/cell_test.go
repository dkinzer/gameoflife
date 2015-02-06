package cell

import (
  "testing"
  p "github.com/dkinzer/gameoflife/point"
)

func TestNewCellStatus(t *testing.T) {
  expected := ALIVE
  actual := Cell{}.Status

  if expected != actual {
    t.Error("New cells are alive.")
  }
}

// What happens to the living.
func TestNextGenerationZeroLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: ALIVE,
    LiveNeighbors: []p.Point{},
  }.nextGeneration()

  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation dies if no living neighbors.")
  }
}

func TestNextGenerationOneLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: ALIVE,
    LiveNeighbors: []p.Point{
      p.Point{0, 1},
    },
  }.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation dies if next to only one living neighbor.")
  }
}

func TestNextGenerationTwoLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: ALIVE,
    LiveNeighbors: []p.Point{
      p.Point{0, 1},
      p.Point{1, 1},
    },
  }.nextGeneration()
  expected := ALIVE
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays alive if next to two living neighbor.")
  }
}

func TestNextGenerationFourLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: ALIVE,
    LiveNeighbors: []p.Point{
      p.Point{0, 1},
      p.Point{1, 1},
      p.Point{-1, 1},
      p.Point{-1, 0},
    },
  }.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation dies if next to four living neighbors.")
  }
}

// What happens to the dead.
func TestNextGenerationThreeLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: DEAD,
    LiveNeighbors: []p.Point{
      p.Point{0, 1},
      p.Point{1, 1},
      p.Point{-1, 1},
    },
  }.nextGeneration()
  expected := ALIVE
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation resurrects if next to three living neighbors.")
  }
}

func TestNextDeadGenerationTwoLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: DEAD,
    LiveNeighbors: []p.Point{
      p.Point{0, 1},
      p.Point{1, 1},
    },
  }.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays dead if next to two living neighbors.")
  }
}

func TestNextDeadGenerationThreeLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: DEAD,
    LiveNeighbors: []p.Point{
      p.Point{0, 1},
      p.Point{1, 1},
      p.Point{-1, 1},
    },
  }.nextGeneration()
  expected := ALIVE
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation resurrects if next to three living neighbors.")
  }
}

func TestNextDeadGenerationFourLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: DEAD,
    LiveNeighbors: []p.Point{
      p.Point{0, 1},
      p.Point{1, 1},
      p.Point{-1, 1},
      p.Point{-1, 0},
    },
  }.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays dead if next to four living neighbors.")
  }
}

func TestNeighborsOfCell(t *testing.T) {
  expected := [p.NEIGHBORS_COUNT_MAX]p.Point{
    p.Point{0, 2},
    p.Point{1, 2},
    p.Point{2, 2},
    p.Point{0, 1},
    p.Point{2, 1},
    p.Point{0, 0},
    p.Point{1, 0},
    p.Point{2, 0},
  }

  actual := Cell{
    Point: p.Point{1, 1},
  }.Neighbors()

  if actual != expected {
    t.Error("Cell.Neighbors()")
  }
}

