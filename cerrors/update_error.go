package cerrors

type UpdateErr struct {
	Err string
}

func (u UpdateErr) Error() string {
	return u.Err
}
