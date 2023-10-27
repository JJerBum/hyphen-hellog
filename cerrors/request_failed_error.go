package cerrors

type RequestFailedErr struct {
	Err string
}

func (r RequestFailedErr) Error() string {
	return r.Err
}
