package helper

func Panicerr(err error) {
	if err != nil {
		panic(err)
	}
}
