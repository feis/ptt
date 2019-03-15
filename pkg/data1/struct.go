package data1

// Board 裡存放一個 Ptt 看板的資料
type Board struct {
	Name     string
	Articles []*Article
}

// Article 裡存放一個 Ptt 文章的資料
type Article struct {
	Title  string
	Author string
	Date   string
	Likes  int
}
