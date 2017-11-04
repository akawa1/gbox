package errr

import (
	"log"
	"testing"
)

func TestNewError(t *testing.T) {
	// e := NewTagCodeError("TEST", 1001, "**%s**", "sth wrong")
	e := NewError("**%s**", "sth wrong")
	log.Print(e.Tag, e.Code)
	log.Println(e.Error())
	log.Println(e.Detail())
	log.Println(e.Stack())
}
