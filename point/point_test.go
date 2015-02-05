package point

import (
  "testing"
)

func TestUpperRightOfPoint(t *testing.T) {
  expected := Point{1, 1}
  actual := Point{0, 0}.upperRight()

  if actual != expected {
    t.Error("Point{0, 0}.upperRight()")
  }
}

func TestUpperCenterOfPoint(t *testing.T) {
  expected := Point{0, 1}
  actual := Point{0, 0}.upperCenter()

  if actual != expected {
    t.Error("Point{0, 0}.upperCenter()")
  }
}

func TestUpperLeftOfPoint(t *testing.T) {
  expected := Point{-1, 1}
  actual := Point{0, 0}.upperLeft()

  if actual != expected {
    t.Error("Point{0, 0}.upperLeft()")
  }
}

func TestCenterLeftOfPoint(t *testing.T) {
  expected := Point{-1, 0}
  actual := Point{0, 0}.centerLeft()

  if actual != expected {
    t.Error("Point{0, 0}.centerLeft()")
  }
}

func TestCenterRightOfPoint(t *testing.T) {
  expected := Point{1, 0}
  actual := Point{0, 0}.centerRight()

  if actual != expected {
    t.Error("Point{0, 0}.centerRight()")
  }
}

func TestBottomRightOfPoint(t *testing.T) {
  expected := Point{1, -1}
  actual := Point{0, 0}.bottomRight()

  if actual != expected {
    t.Error("Point{0, 0}.bottomRight()")
  }
}

func TestBottomCenterOfPoint(t *testing.T) {
  expected := Point{0, -1}
  actual := Point{0, 0}.bottomCenter()

  if actual != expected {
    t.Error("Point{0, 0}.bottomCenter()")
  }
}

func TestBottomLeftOfPoint(t *testing.T) {
  expected := Point{-1, -1}
  actual := Point{0, 0}.bottomLeft()

  if actual != expected {
    t.Error("Point{0, 0}.bottomLeftFrom")
  }
}

func TestNeighborsOfPoint(t *testing.T) {
  expected := [NEIGHBORS_COUNT_MAX]Point{
    Point{-1, 1},
    Point{0, 1},
    Point{1, 1},
    Point{-1, 0},
    Point{1, 0},
    Point{-1, -1},
    Point{0, -1},
    Point{1, -1},
  }

  actual := Point{0,0}.neighbors()

  if actual != expected {
    t.Error("NeighborsOfPoint(Point{0, 0}")
  }
}

