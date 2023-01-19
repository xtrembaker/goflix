package movie

type Title struct {
	string
}

func (t Title) GetValue() string {
	return t.string
}

func NewTitle(title string) Title {
	if len(title) <= 0 {
		panic("Title must have at least one character")
	}

	return Title{title}
}
