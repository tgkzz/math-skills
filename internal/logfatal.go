package internal

import "log"

func Logfatal(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
