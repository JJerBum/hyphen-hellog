package cerrors

type SelectErr struct {
	Err string
}

func (s SelectErr) Error() string {
	return s.Err
}
