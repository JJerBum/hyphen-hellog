package cerrors

type CreateErr struct {
	Err string
}

func (c CreateErr) Error() string {
	return c.Err
}
