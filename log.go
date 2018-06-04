package log

import (
	"fmt"
	"time"
)

func Writeconsole() {

	var sModo string

	switch sTipo {
	case "W":
		sModo = "WARNING!!! "
	case "E":
		sModo = "** ERROR **"
	case "T":
		sModo = "** TEXTO **"
	default:
		sModo = "           "
	}

	fmt.Printf("%s - %s - %s - %s \n", time.Now().Format("2006/01/02 - 03:04:05 PM"), sModo, sProcesso, sMsg)

}
