package model

// Data msg will be push to databus
type Data struct {
	Action string `json:"action"`
	Table  string `json:"table"`
	// kafka key
	Key string                 `json:"-"`
	Old map[string]interface{} `json:"old,omitempty"`
	New map[string]interface{} `json:"new,omitempty"`
}