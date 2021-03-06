package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestJson(t *testing.T) {

	bolB, _ := json.Marshal(true)
	assert.Equal(t, "true", string(bolB))

	intB, _ := json.Marshal(1)
	assert.Equal(t, "1", string(intB))

	fltB, _ := json.Marshal(1.23)
	assert.Equal(t, "1.23", string(fltB))

	strB, _ := json.Marshal("gopher")
	assert.Equal(t, "\"gopher\"", string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	assert.Equal(t, "[\"apple\",\"peach\",\"pear\"]", string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	assert.Equal(t, "{\"apple\":5,\"lettuce\":7}", string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	expected := "{\"Page\":1,\"Fruits\":[\"apple\",\"peach\",\"pear\"]}"
	assert.Equal(t, expected, string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	expected = "{\"page\":1,\"fruits\":[\"apple\",\"peach\",\"pear\"]}"
	assert.Equal(t, expected, string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	exp := make(map[string]interface{})
	exp["num"] = 6.13
	exp["strs"] = []interface{}{"a", "b"}
	assert.Equal(t, exp, dat)

	num := dat["num"].(float64)
	assert.Equal(t, 6.13, num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	assert.Equal(t, "a", str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	assert.Equal(t, "apple", res.Fruits[0])

	// Not exactly sure how to test this since it's printing to stdout
	// enc := json.NewEncoder(os.Stdout)
	// d := map[string]int{"apple": 5, "lettuce": 7}
	// enc.Encode(d)
}
