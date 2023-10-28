package cerrors

type WrongApproachErr struct {
	Err string
}

func (w WrongApproachErr) Error() string {
	return w.Err
}
