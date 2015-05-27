package raygun4go

import "log"

// OutputHandler is the type of function for handling output. It receives a string,
// handles it and returns nil if ok, an error if something went wrong.
type OutputHandler func(string) error

// OutputLog is an OutputHandler that logs the given message
var OutputLog OutputHandler = func(message string) error {
	log.Println(message)
	return nil
}

// OutputMute is a OutputHandler that does nothing with the given message
var OutputMute OutputHandler = func(message string) error {
	return nil
}
