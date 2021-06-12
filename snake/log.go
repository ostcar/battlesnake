package snake

import "log"

const debug = false

func debugLog(format string, a ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, a...)
}
