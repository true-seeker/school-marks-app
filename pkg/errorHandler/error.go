package errorHandler

import "log"

type WebError struct {
	Err  error
	Code int
}

// FailOnError Аварийно завершить программу и вывести ошибку
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
