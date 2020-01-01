package main

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
}

func TestXml(t *testing.T) {
	assert.Equal(t, 1, 1)
}
