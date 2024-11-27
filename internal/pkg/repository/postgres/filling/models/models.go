package models

import "fmt"

/*
	Для дефолтного заполнения БД
*/

// Subscription Структура для подписки
type Subscription struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Layer       int    `json:"layer"`
}

// Post Структура для постов
type Post struct {
	Title         string `json:"title"`
	Info          string `json:"info"`
	MediaFilename string `json:"mediaFilename"`
	Layer         int    `json:"layer"`
}

// FillingAuthor Основная структура для каждого автора
type FillingAuthor struct {
	AuthorName         string         `json:"authorName"`
	Password           string         `json:"password"`
	CustomSubscription []Subscription `json:"customSubscription"`
	Posts              []Post         `json:"posts"`
}

func (s *Subscription) String() string {
	ln := 3
	rTitle := []rune(s.Title)
	title := rTitle[:min(ln, len(rTitle))]
	rDescription := []rune(s.Description)
	description := rDescription[:min(ln, len(rDescription))]
	return fmt.Sprintf("Sub{ title: %v..., desc: %v..., cost: %v, layer: %v }",
		title, description, s.Cost, s.Layer)
}

func (p *Post) String() string {
	ln := 3
	rTitle := []rune(p.Title)
	title := rTitle[:min(ln, len(rTitle))]
	rInfo := []rune(p.Info)
	info := rInfo[:min(ln, len(rInfo))]
	return fmt.Sprintf("Post {title: %v..., info: %v..., mediaFilename: %v, layer: %v}",
		title, info, p.MediaFilename, p.Layer)
}

func (f *FillingAuthor) String() string {
	subs := ""
	for _, sub := range f.CustomSubscription {
		subs += sub.String() + ", \n\t\t"
	}
	posts := ""
	for _, post := range f.Posts {
		posts += post.String() + ", \n\t\t"
	}
	return fmt.Sprintf("\n\nFAuthor{ \n\tAuthorName: %v, \n\tPassword: %v, \n\tCustomSub: [%v], \n\tPosts: [%v] }", f.AuthorName,
		f.Password, subs, posts)
}
