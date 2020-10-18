package gowait

import "encoding/json"

type Taskdef struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Upstream string  `json:"upstream,omitempty"`
	Parent   *string `json:"parent"`
	Meta     Dict    `json:"meta"`

	Inputs json.RawMessage `json:"inputs"`
}

func (t *Taskdef) Input(out interface{}) error {
	return json.Unmarshal(t.Inputs, out)
}
