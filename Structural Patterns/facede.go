package main

type FacedeGeteway struct {
	ArticleAPI ArticleAPI
	MusicAPI MusicAPI
}

func (F *FacedeGeteway) CreateArticles() *ArticleSrv {
	data := F.CreateArticle(1,"2","123")
	return data
}

func (F *FacedeGeteway) CreateMusics() *MusicSrv {
	data := F.CreateMusic("A",1,"A")
	return data
}

func (F *FacedeGeteway) CreateArticle(ArticleID int, ArticleTitle string, content string) *ArticleSrv  {
	data := F.ArticleAPI.Create(ArticleID, ArticleTitle, content)
	return data
}

func (F *FacedeGeteway) CreateMusic(MusicType string, MusicID int, MusicName string) *MusicSrv {
	data := F.MusicAPI.Create(MusicType, MusicID, MusicName)
	return data
}

type ArticleAPI struct {}

type ArticleSrv struct {
	ArticleID int
	ArticleTitle string
	content string
}

func (A *ArticleAPI) Create(ArticleID int, ArticleTitle string, content string) *ArticleSrv {
	return &ArticleSrv{
		ArticleID:  ArticleID  ,
		ArticleTitle: ArticleTitle,
		content:      content,
	}
}

type MusicAPI struct {}

type MusicSrv struct {
	MusicType string
	MusicID int
	MusicName string
}

func (M *MusicAPI) Create(MusicType string, MusicID int, MusicName string) *MusicSrv {
	return &MusicSrv{
		MusicType: MusicType,
		MusicID:   MusicID,
		MusicName: MusicName,
	}
}

func main() {
	f := FacedeGeteway{}
	f.CreateArticles()
	f.CreateMusics()
}