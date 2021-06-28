package util

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func CheckErrorRetVal(val interface{}, err error) interface{} {
	CheckError(err)
	return val
}