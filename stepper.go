// Stepper package provides super simple steps handler and
// formatter.
//
// Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
package stepper

import (
	"fmt"
	"os"
)

// Possible steps
const (
	done = iota
	busy
	fail
)

// Stepper contains information about current step's
// state and provides functions to operate on it.
type Stepper struct {
	step string
}

// doStep performs transition to desired step.
func (s *Stepper) doStep(status int) {
	var sstr string
	switch status {
	case done:
		sstr = "\033[32mdone\033[0m"
	case busy:
		sstr = "\033[34mbusy\033[0m"
	case fail:
		sstr = "\033[31mfail\033[0m"
	}
	fmt.Fprintf(os.Stderr, "\r%s... %s", s.step, sstr)
}

// Fail terminates currently performed step (if any) with
// a failure notification. If second param is true, then
// it exits the program as well.
func (s *Stepper) Fail(msg string, exit bool) {
	if s.step != "" {
		s.doStep(fail)
		fmt.Fprintf(os.Stderr, "\n\033[31m!!! %s\033[0m\n", msg)
		if exit {
			os.Exit(1)
		}
	}
}

// Ok terminates currently performed step (if any) with
// a success notification.
func (s *Stepper) Ok() {
	if s.step != "" {
		s.doStep(done)
		s.step = ""
		println()
	}
}

// do setups step for specified message, executes passed operation
// and displays appropriate results.
func (s *Stepper) do(msg string, exit bool, fn func()(error)) {
	s.step = msg
	s.doStep(busy)
        if err := fn(); err != nil {
		s.Fail(err.Error(), exit)
		return
	}
	s.Ok()
}

// MustDo executes operation and exists when error occurs.
func (s *Stepper) MustDo(msg string, fn func()(error)) {
	s.do(msg, true, fn)
}

// Do executes operation and lets program to keep going when
// error occurs.
func (s *Stepper) Do(msg string, fn func()(error)) {
	s.do(msg, false, fn)
}
