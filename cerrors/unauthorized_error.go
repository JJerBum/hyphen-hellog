package cerrors

type UnauthorizedErr struct {
	Err string
}

func (u UnauthorizedErr) Error() string {
	return u.Err
}
