package sqlite3

type ErrNoExtended int

const (
	ErrConstraint           = 1
	ErrConstraintUnique     = ErrNoExtended(2)
	ErrConstraintForeignKey = ErrNoExtended(3)
)

type Error struct {
	Code         int
	ExtendedCode ErrNoExtended
}

func (err Error) Error() string {
	return ""
}
