package gameoflife

type status byte

const (
  EMPTY status = iota
  DEAD
  OCCUPIED
  ALIVE
)

type Cell struct {
  Point
  Status status
  Neighbors []Point
}

