# Go Stepper

Stepper package provides super simple steps handler and
formatter.

## Installation

Install with `go get` tool:

	$ go get github.com/nu7hatch/gostepper

## Usage

Use `Do` and `MustDo` functions to perform operations. If error
occurs then `Do` lets program to keep being executed, while `MustDo`
exists the program. Here's an example:

	package main
	
	import stepper "github.com/nu7hatch/gostepper"

	func main() {
	    var s stepper.Stepper

	    s.Do("Configuring something", func() error {
            return ConfigureSomething()
        })
        s.MustDo("Starting stomething else", func() error {
            return StartSomething()
        })
	}

## Copyright

Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>

See LICENSE file for details.
