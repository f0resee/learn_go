package gofileoperation

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func TestJsonGroup(t *testing.T) {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	var groups []ColorGroup
	groups = append(groups,group)
	groups = append(groups,group)
	b, err := json.Marshal(groups)
	if err != nil {
		t.Log("error", err)
	}
	t.Log(string(b))
}
func TestJsonRead(t *testing.T) {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		t.Log("error", err)
	}
	t.Log(string(b))

	file, errw := os.OpenFile("test.json", syscall.O_RDWR, 0666)
	if errw != nil {
		t.Log("Open test.json failed")
	}
	file.Write(b)
	file.Close()
}

func TestJsonWrite(t *testing.T) {
	jsonb := []byte(
		`[{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}]`)
	var groups []ColorGroup
	err := json.Unmarshal(jsonb, &groups)
	if err != nil {
		t.Log("Decode failed")
	}
	t.Log(groups)

	file, errr := os.Open("test.json")
	if errr != nil {
		t.Log("open test.json failed")
	}
	jb, er := ioutil.ReadAll(file)
	if er != nil {
		t.Log("Read failed")
	}
	t.Log(string(jb))
	var obj ColorGroup
	e := json.Unmarshal(jb, &obj)
	if e != nil {
		t.Log("Unmarshal failed")
	}
	t.Log(obj)
	file.Close()

}
