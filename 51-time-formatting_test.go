package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormatting(t *testing.T) {
	assert.Equal(t, 1, 1)

	now := time.Date(
		2020, 01, 01, 01, 01, 01, 651387237, time.UTC)
	assert.Equal(t, "2020-01-01T01:01:01Z", now.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	assert.Equal(t, "2012-11-01 22:08:41 +0000 +0000", t1.String())
	fmt.Println(now.Format("3:04PM"))
	assert.Equal(t, "1:01AM", now.Format("3:04PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(now.Format("Mon Jan _2 2006 15:04:05"))
	fmt.Println(now.Format("2006-01-02-T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	fmt.Println(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	fmt.Println(e)
}
