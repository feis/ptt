package data

// NewBoard 會建構一個新的看板
func NewBoard(name string) *Board {
	return &Board{
		name: name,
	}
}
