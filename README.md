# SIE Message Module for go-nmsg

This module provides definitions for message types from [sie-nmsg](https://github.com/farsightsec/sie-nmsg)
for use with the [go-nmsg](https://github.com/farsightsec/go-nmsg) library.

## Synopsis

Importing the module with:

	import nmsg_sie "github.com/farsightsec/go-nmsg_sie"

will register several `sie_nmsg` types with the `go-nmsg` module, allowing them
to be passed to `nmsg.Payload()` and returned from the `NmsgPayload` method `Message()` 


## Usage Example

	import "github.com/farsightsec/go-nmsg"
	import "github.com/farsightsec/go-nmsg_sie"

	...
	input := nmsg.NewInput(r, mtu)
	var countA, countAAAA uint64
	for {
		payload, err := input.Recv()
		if (err != nil) {
			if (nmsg.IsDataError(err)) {
				continue
			}
			break
		}

		message := payload.Message()
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

	fmt.Printf("%d A records, %d AAAA records\n", countA, countAAAA)
