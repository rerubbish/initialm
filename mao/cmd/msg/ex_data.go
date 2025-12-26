package msg

type ExData struct {
	Name string `json:"name"`
	Data map[string]struct {
		Type  Type `json:"type"`
		Value any    `json:"value"`
	} `json:"data"`
}

// type Type string
type Type string

const (
	TypeText Type = "text"
)
