package style

type Style struct {
	Question     string
	Placeholder  string
	Cursor       string
	Selected     string
	Unselected   string
	ErrorColor   string
	SuccessColor string
}

func Default() *Style {
	return &Style{
		Question:     "\033[1m\033[36m?\033[0m",
		Placeholder:  "\033[90m%s\033[0m",
		Cursor:       "\033[36m❯\033[0m",
		Selected:     "\033[32m●\033[0m",
		Unselected:   "\033[90m○\033[0m",
		ErrorColor:   "\033[31m",
		SuccessColor: "\033[32m",
	}
}

func NewStyle(opts ...StyleOption) *Style {
	s := Default()
	for _, opt := range opts {
		opt(s)
	}
	return s
}

type StyleOption func(*Style)

func WithQuestion(q string) StyleOption {
	return func(s *Style) {
		s.Question = q
	}
}
