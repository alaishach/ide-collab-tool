package panics

func PanicErr(msg string, err error) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}

func PanicDB(funcName string, err error) {
	if err != nil {
		panic("Error with db function '" + funcName + "': " + err.Error())
	}
}
