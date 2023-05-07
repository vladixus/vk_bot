package models

type Keyboard struct {
	OneTime bool    `json:"one_time"`
	Buttons Buttons `json:"buttons"`
}

type Buttons [][]struct {
	Action Action `json:"action"`
	Color  string `json:"color"`
}

type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Label   string `json:"label"`
}
