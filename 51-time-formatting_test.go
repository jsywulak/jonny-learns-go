package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormatting(t *testing.T) {
	now := time.Date(
		2020, 01, 01, 01, 01, 01, 651387237, time.UTC)
	assert.Equal(t, "2020-01-01T01:01:01Z", now.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	assert.Equal(t, "2012-11-01 22:08:41 +0000 +0000", t1.String())
	assert.Equal(t, "1:01AM", now.Format("3:04PM"))
	assert.Equal(t, "Wed Jan  1 01:01:01 2020", now.Format("Mon Jan _2 15:04:05 2006"))
	assert.Equal(t, "Wed Jan  1 2020 01:01:01", now.Format("Mon Jan _2 2006 15:04:05"))
	assert.Equal(t, "2020-01-01-T01:01:01.651387+00:00", now.Format("2006-01-02-T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	assert.Equal(t, "0000-01-01 20:41:00 +0000 UTC", t2.String())

	actual := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	assert.Equal(t, "2020-01-01T01:01:01-00:00", actual)

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	assert.Equal(t, `parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"`, e.Error())
}
