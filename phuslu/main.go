package main

// sample application for using the github.com/phuslu/log
// log framework project
import (
	"github.com/phuslu/log"
)

func main() {
	log.DefaultLogger = log.Logger{
		Level:      log.DebugLevel,
		Caller:     1,
		TimeField:  "date",
		TimeFormat: "2006-01-02",
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
		},
	}
	log.Debug().Msg("This is a debug")
	log.Info().Str("foo", "bar").Msgf("This will have key=val printed")
	log.Warn().Msg("This is a warning message")
	log.Warn().Bool("booleanKey",false).Msgf("This will have key=val printed (boolean)")
	log.Printf("Hello, %s  this is just a normal printf", "Human")
	anotherFunction()
}

func anotherFunction(){
	log.Info().Stack(true).Msg("This will print the whole stack call")
}
