package models

type NanoTimestamp = float64

type SenML struct {
	N string        `json:"n"` // Resource Name
	T NanoTimestamp `json:"t"` // Timestamp
	V float32       `json:"v"` // Value
	U string        `json:"u"` // Unit
	L string        `json:"l"` // Link
}

type BaseSenML struct {
	N string  `json:"n,omitempty"` // Resource Name
	V float32 `json:"v"`           // Value
	U string  `json:"u,omitempty"` // Unit

	BN   string        `json:"bn"`           // Prefix name
	BT   NanoTimestamp `json:"bt"`           // Base timestamp
	BU   string        `json:"bu,omitempty"` // Base unit
	BVER int           `json:"bver"`         // Base version
	L    string        `json:"l"`            // Link
}
