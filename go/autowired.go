package main

import (
	"fmt"
	"time"
)

// sign.go
type signer struct {
	now func() time.Time
}

func (s *signer) SignBetter() time.Time {
	return s.now()
}

func newSigner() *signer {
	return &signer{
		now: time.Now,
	}
}

var _timeNow = time.Now

func signbad() time.Time {
	now := _timeNow()
	return now
}
func main() {
	// bad
	fmt.Println(signbad())
	fmt.Println(newSigner().SignBetter())

}
