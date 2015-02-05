package gameoflife

import (
  "testing"
)

func TestNewCellStatus(t *testing.T) {
  expected := ALIVE
  actual := Cell{}.Status

  if expected != actual {
    t.Error("New cells are alive.")
  }
}

func TestNextGenerationZeroLiveNeighbors(t *testing.T) {
  cell := Cell{}.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation is dies if no live neighbors.")
  }
}

func TestNextGenerationOneLiveNeighbors(t *testing.T) {
  cell := Cell{
    LiveNeighbors: []Point{
      Point{0, 1},
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
    LiveNeighbors: []Point{
      Point{0, 1},
      Point{1, 1},
    },
  }.nextGeneration()
  expected := ALIVE
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays alive if next to two living neighbor.")
  }
}

func TestNextGenerationThreeLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: DEAD,
    LiveNeighbors: []Point{
      Point{0, 1},
      Point{1, 1},
      Point{-1, 1},
    },
  }.nextGeneration()
  expected := ALIVE
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays alive if next to three living neighbors.")
  }
}

func TestNextDeadGenerationTwoLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: DEAD,
    LiveNeighbors: []Point{
      Point{0, 1},
      Point{1, 1},
    },
  }.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays dead if next to two living neighbor.")
  }
}

func TestNextDeadGenerationThreeLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: DEAD,
    LiveNeighbors: []Point{
      Point{0, 1},
      Point{1, 1},
      Point{-1, 1},
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
    LiveNeighbors: []Point{
      Point{0, 1},
      Point{1, 1},
      Point{-1, 1},
      Point{-1, 0},
    },
  }.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays dead if next to four living neighbors.")
  }
}

func TestNextGenerationFourLiveNeighbors(t *testing.T) {
  cell := Cell{
    Status: ALIVE,
    LiveNeighbors: []Point{
      Point{0, 1},
      Point{1, 1},
      Point{-1, 1},
      Point{-1, 0},
    },
  }.nextGeneration()
  expected := DEAD
  actual := cell.Status

  if expected != actual {
    t.Error("Next generation stays dies if next to four living neighbors.")
  }
}
