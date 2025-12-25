package msg

type ExData struct {
	Name string `json:"name"`
	Data map[string]struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	} `json:"data"`
}
