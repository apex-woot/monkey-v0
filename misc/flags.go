package misc

import (
	"flag"
	"io"
	"log"
)

const VERBOSE_LOGGING = "vLog"

func ParseAndApplyFlags() {
	verboseLogging := flag.Bool(VERBOSE_LOGGING, false, "Enable verbose lexer logging")
	flag.Parse()
	if !*verboseLogging {
		log.SetOutput(io.Discard)
	}
}
