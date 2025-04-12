package assert

import (
	"log"
)

// TODO: Add some debug logs?

func Assert(truth bool, message string) {
	if !truth {
		log.Fatal(message)
	}
}
