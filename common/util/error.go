package util

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
