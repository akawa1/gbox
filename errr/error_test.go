package errr

import (
	"log"
	"testing"
)

func TestNewError(t *testing.T) {
	e := NewErrorWithTag("TEST", "sth wrong")
	log.Println(e.Error())
	log.Println(e.Detail())
	log.Println(e.Stack())
}
