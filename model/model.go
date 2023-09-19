package model

import (
	"encoding/json"
	"log"
)

type Student struct {
	FirstName string
	LastName  string `json:"ho"`
	Age       int    `json:"tuoi"`
	Grade     float32
	ClassName string
}

func (b Student) ReceiverGetFullName() string {
	return b.FirstName + " " + b.LastName

}
func (b Student) ReceiverGetFirstName() string {
	return b.FirstName

}

func (c *Student) SetName(a, b string) {
	c.FirstName = a
	c.LastName = b

}

func (c *Student) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}

func (c *Student) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}

func FromJsonNormal(a string, c *Student) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}

}
