package data

// Article 是存放文章資訊的結構
type Article struct {
	Title  string
	Author string
	Date   string
	Likes  int
}

// Board 是存放看板資訊的結構
type Board struct {
	Name     string
	Articles []Article
}
