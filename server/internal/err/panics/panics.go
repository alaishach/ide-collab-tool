package panics

func PanicErr(msg string, err error) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}

func PanicDB(funcName string, err error) {
	if err != nil {
		panic("Error with postgresql db function '" + funcName + "': " + err.Error())
	}
}

func PanicRedis(funcName string, err error) {
	if err != nil {
		panic("Error with postgresql db function '" + funcName + "': " + err.Error())
	}
}

func PanicMisuse(funcName string, description string) {
	panic("Misuse of function" + funcName + ": " + description)
}
