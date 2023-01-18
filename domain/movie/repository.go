package movie

type Repository interface {
	List() []*Movie
}
