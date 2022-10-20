package errs

import "log"

func Panic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
