package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

type Book struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

func TestJson(t *testing.T) {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	t.Log("\n")
}

func TestJsonRawMessage(t *testing.T) {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, c := range colors {
		var dst any
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}

func TestJsonUnmashall(t *testing.T) {
	book := &Book{
		Title: "Chinese",
		Price: 10,
	}
	data, err := json.Marshal(book)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(data))
	str := "{\"title\":\"Chinese\",\"price\":10}"
	var book2 Book
	err1 := json.Unmarshal([]byte(str), &book2)
	if err1 != nil {
		t.Log(err1)
	}
	t.Log(book2)
}
