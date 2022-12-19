package response

type Book struct {
	ID         uint   `json:"id" gorm:"primary_key" example:"1"`
	Title      string `json:"title" example:"This is Title"`
	CoverImage string `json:"cover_image" example:"http://cdn.image.com/path/to/image.png"`
	Category   string `json:"category" example:"blade"`
}

type BookDetail struct {
	ID         uint      `json:"id" gorm:"primary_key" example:"1"`
	Title      string    `json:"title" example:"This is Title"`
	Summary    string    `json:"summary" example:"test lipsum"`
	CoverImage string    `json:"cover_image" example:"http://cdn.image.com/path/to/image.png"`
	Category   string    `json:"category" example:"blade"`
	Chapter    []Chapter `json:"chapters"`
}

type Chapter struct {
	ID           string `json:"id" example:"1"`
	Title        string `json:"title" example:"blaze"`
	MediaPath    string `json:"media_path" example:"http://cdn.audio.com/path/to/audio.wav"`
	BookDetailID string `json:"book_id" example:"1"`
}
