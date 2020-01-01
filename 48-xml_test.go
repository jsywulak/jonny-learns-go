package main

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.ID, p.Name, p.Origin)
}

func TestXml(t *testing.T) {
	assert.Equal(t, 1, 1)

	coffee := &Plant{ID: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}

}
