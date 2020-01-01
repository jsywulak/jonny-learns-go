package main

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func TestXml(t *testing.T) {
	assert.Equal(t, 1, 1)
}
