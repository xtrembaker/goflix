package movie

type Repository interface {
	List() []*Movie
	Get(id int64) (*Movie, error)
	Save(*Movie)
}
