package handlers

func ThrowError(err error) {
	if err != nil {
		panic(err) // App Crash
	}
}
