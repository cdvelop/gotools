package gotools

import "fmt"

// opt: ok,err,warn,info
func Print(opt string, message interface{}) {
	var color string
	switch opt {
	case "ok":
		color = "32" //green
	case "err":
		color = "31" //red
	case "warn":
		color = "33" //yellow
	case "info":
		color = "36" //magenta blue=34
	default:
		color = "0"
	}

	fmt.Printf("\033[%sm%s\033[0m", color, message)
}
