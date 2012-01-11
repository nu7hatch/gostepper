// Stepper package provides super simple steps handler and
// formatter.
//
// Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
package stepper

import (
	"fmt"
	"strings"
	"os"
)

// Stepper contains information about current step's
// state and provides functions to operate on it.
type Stepper struct {
	step string
}

func (s *Stepper) doStep(status byte) {
	var sstr string
	switch status {
	case 0: // DONE
		sstr = "\033[32mdone\033[0m"
	case 1: // BUSY
		sstr = "\033[34mbusy\033[0m"
	case 2: // FAIL
		sstr = "\033[31mfail\033[0m"
	}
	sep := 80 - len(s.step) - 6
	fmt.Printf("\r%s %s %s", s.step, strings.Repeat(".", sep), sstr)
}

// Fail terminates currently performed step (if any) with
// a failure notification. If second param is true, then
// it exits the program as well.
func (s *Stepper) Fail(msg string, exit bool) {
	if s.step != "" {
		s.doStep(2)
		fmt.Printf("\n\033[31m!!! %s\033[0m\n", msg)
		if exit {
			os.Exit(1)
		}
	}
}

// Ok terminates currently performed step (if any) with
// a success notification.
func (s *Stepper) Ok() {
	if s.step != "" {
		s.doStep(0)
		s.step = ""
		println()
	}
}

// Start setups step for specified message.
func (s *Stepper) Start(msg string, params ...interface{}) {
	s.step = fmt.Sprintf(msg, params...)
	s.doStep(1)
}