package cerrors

type ValidationErr struct {
	Err string
}

func (v ValidationErr) Error() string {
	return v.Err
}
