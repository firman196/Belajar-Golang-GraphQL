package utils

import "log"

func SetFatalError(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}
