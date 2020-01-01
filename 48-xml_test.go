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
	var actual string

	coffee := &Plant{ID: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	actual = string(out)
	expected := ` <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>`
	assert.Equal(t, expected, actual)

	actual = (xml.Header + string(out))
	expected = `<?xml version="1.0" encoding="UTF-8"?>
 <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>`
	assert.Equal(t, expected, actual)

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	expected = "Plant id=27, name=Coffee, origin=[Ethiopia Brazil]"
	actual = p.String()
	assert.Equal(t, expected, actual)

	tomato := &Plant{ID: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
	actual = string(out)
	expected = ` <nesting>
   <parent>
     <child>
	   <plant id="27">
	     <name>Coffee</name>
	     <origin>Ethiopia</origin>
	     <origin>Brazil</origin>
	   </plant>
	   <plant id="81">
         <name>Tomato</name>
         <origin>Mexico</origin>
         <origin>California</origin>
       </plant>
     </child>
   </parent>
 </nesting>`

}
