package jarvis

import (
	"strings"
)


/*
	Continuously listens to user input. Possibly keeps history for 
	context i.e. "long term / short term memory" of conversation.
*/
func ParseInput(inputChannel chan string, quit chan struct{}) {
	var input string
	for {
		select {
			case input = <- inputChannel:
				input = strings.Trim(strings.ToLower(input), " ")
				println(input)
			case <- quit:
				return
		}
	}
}