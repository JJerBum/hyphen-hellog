package cerrors

type ParsingErr struct {
	Err string
}

func (p ParsingErr) Error() string {
	return p.Err
}
