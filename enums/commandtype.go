package enums

type CommandType int

const (
	M CommandType = iota
	R
	L
)

func (c CommandType) String() string {
	return [...]string{"M", "R", "L"}[c]
}
