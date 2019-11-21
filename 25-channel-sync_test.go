package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannelSync(t *testing.T) {
	assert.Equal(t, 1, 1)
	dummymain()
}

func dummymain() {

}

func worker(done chan bool) {
	fmt.Print("working")
	time.Sleep(time.Second)
	fmt.Println("done")
}
