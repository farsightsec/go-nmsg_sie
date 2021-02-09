# SIE Message Module for go-nmsg

This module provides definitions for message types from [sie-nmsg](https://github.com/farsightsec/sie-nmsg)
for use with the [go-nmsg](https://github.com/farsightsec/go-nmsg) library.

## Synopsis

Importing the module with:

	import nmsg_sie "github.com/farsightsec/go-nmsg_sie"

will register several `nmsg_sie` types with the `go-nmsg` module, allowing them
to be passed to `nmsg.Payload()` and returned from the `NmsgPayload` method `Message()` 


## Usage Example
	
	package main
	
	import (
		"os"
		"fmt"
		"github.com/farsightsec/go-nmsg"
		"github.com/farsightsec/go-nmsg_sie"
	)
	
	func main() {
		var countA, countAAAA uint64
		input := nmsg.NewInput(os.Stdin, nmsg.MaxContainerSize)
		for {
			payload,err := input.Recv()
			if err != nil {
				break
			}
			message, err := payload.Message()
			if err != nil {
				continue
			}
			m, ok := message.(*nmsg_sie.DnsDedupe)
			if !ok {
				continue
			}
			switch m.GetRrtype() {
			case 1:
				countA += uint64(m.GetCount())
			case 28:
				countAAAA += uint64(m.GetCount())
			}
		}
		fmt.Printf("%d A, %d AAAA\n", countA, countAAAA)
	}
