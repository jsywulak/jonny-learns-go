package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormatting(t *testing.T) {
	assert.Equal(t, 1, 1)

	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println(t1)

	fmt.Println(now.Format("3:04PM"))
}
