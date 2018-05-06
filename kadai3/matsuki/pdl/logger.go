package pdl

import "log"

func debug(msg string) {
	log.Print("[DEBUG] " + msg)
}

func warn(msg string) {
	log.Print("[WARN] " + msg)
}

// errorであってもerrであっても被ることが多いのでここだけはerrorlogにする
func errorlog(msg string) {
	log.Print("[ERROR] " + msg)
}
