package cerrors

type DeleteErr struct {
	Err string
}

func (d DeleteErr) Error() string {
	return d.Err
}
