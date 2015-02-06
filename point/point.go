package point

const NEIGHBORS_COUNT_MAX = 8
const NEIGHBORS_COUNT_MIN = 0

type Point struct {
  X int
  Y int
}

type Pointer interface {
  UpperRight() Point
}


func (p Point) Neighbors() [NEIGHBORS_COUNT_MAX]Point {
  return [NEIGHBORS_COUNT_MAX]Point{
    p.upperLeft(),
    p.upperCenter(),
    p.upperRight(),
    p.centerLeft(),
    p.centerRight(),
    p.bottomLeft(),
    p.bottomCenter(),
    p.bottomRight(),
  }
}

func (p Point) upperLeft() Point {
  return Point{p.X - 1, p.Y + 1}
}

func (p Point) upperCenter() Point {
  return Point{p.X, p.Y + 1}
}

func (p Point) upperRight() Point {
  return Point{p.X + 1, p.Y + 1}
}

func (p Point) centerLeft() Point {
  return Point{p.X - 1, p.Y}
}

func (p Point) centerRight() Point {
  return Point{p.X + 1, p.Y}
}

func (p Point) bottomLeft() Point {
  return Point{p.X - 1, p.X - 1}
}

func (p Point) bottomCenter() Point {
  return Point{p.X, p.X - 1}
}

func (p Point) bottomRight() Point {
  return Point{p.X + 1, p.X - 1}
}

