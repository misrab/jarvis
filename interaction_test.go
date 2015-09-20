package jarvis


import (
	"testing"

)

var (
	SALUTATIONS = []string{
		"hi",
		"hello",
	}
)


func TestParseInputSalutation(t *testing.T) {
	inputChannel := make(chan string)
	quit := make(chan struct{})

	go ParseInput(inputChannel, quit)

	for _, salutation := range SALUTATIONS {
		// println(salutation)
		inputChannel <- salutation
	}

	quit <- struct{}{}

}