package repository

type Repository struct {
	Auth     AuthRepository
	User     UserRepository
	Book     BookRepository
	Chapter  ChapterRepository
	Category CategoryRepository
}
