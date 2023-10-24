package exception

func Sniff(err error) {
	if err != nil {
		panic(err)
	}
}
