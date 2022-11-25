package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Student struct {
	Name string    `json:"name"`
	ID   uint      `json:"id,omitempty"`
	Time time.Time `json:"-"`
}

func TestJson(t *testing.T) {
	s := Student{
		Name: "bean",
		ID:   2021140937,
		Time: time.Now(),
	}
	marshal, err := json.Marshal(s)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))

	ss := Student{}
	err = json.Unmarshal(marshal, &ss)
	if err != nil {
		return
	}
	fmt.Println(ss)
}
