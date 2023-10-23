package model

import (
	"encoding/json"
	"log"
	"time"
)

type Student struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age,omitempty”`
	Grade        float32   `json:"grade,omitempty”`
	ClassName    string    `json:"class_name,omitemty"`
	EntranceDate time.Time `json:"entrance_date"`
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
