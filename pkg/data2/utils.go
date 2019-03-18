package data2

// NewBoard 會產生一個名稱為 n 的看板資料
func NewBoard(n string) *Board {
	return &Board{
		name: n,
	}
}
