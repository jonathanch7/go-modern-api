package logger

import "log"

func Command(cmdName string, cmdValue interface{}, err error) {
	if err != nil {
		log.Printf("%s command error %v cmd: %+v\n", cmdName, err, cmdValue)
	} else {
		log.Printf("%s command succeeded cmd: %+v\n", cmdName, cmdValue)
	}
}
