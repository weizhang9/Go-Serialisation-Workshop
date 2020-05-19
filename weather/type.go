package weather

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Unit string

const (
	Celcius    Unit = "c"
	Fahrenheit Unit = "f"
	Kelvin     Unit = "k"
	Inch       Unit = "in"
	Centimeter Unit = "cm"
)

// Value if a measurement value
type Value struct {
	Value float64 `json:"value"`
	Unit  Unit    `json:"unit"`
}

// MarshalJSON implement the json.Marshaller interface
// to change the Marshal behaviour
// here it acts as a wrapper
func (v Value) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%f%s", v.Value, v.Unit)
	return json.Marshal(s)
}

// UnmarshalJSON allows to unmarshal customised marshalled JSON data
func (v *Value) UnmarshalJSON(data []byte) error {
	// data example "48.200000f"
	s := string(data[1 : len(data)-1]) // trim enclosing quotation marks
	i := strings.LastIndexAny(s, "0123456789")
	if i == -1 {
		return fmt.Errorf("no number in %#v", string(data))
	}

	if i == len(s)-1 {
		return fmt.Errorf("no unit in %#v", string(data))
	}

	i++ // move to unit
	val, err := strconv.ParseFloat(s[:i], 64)
	if err != nil {
		return err
	}
	v.Value = val
	v.Unit = Unit(s[i:])
	return nil
}

// Record of measurement
type Record struct {
	Time        time.Time `json:"time"`
	Station     string    `json:"station"`
	Temperature Value     `json:"temperature"`
	Rain        Value     `json:"rain"`
}

/* If you need different format for serialising time
encoding/json package use RFC3339 format
JTime embeds time.Time for custom JSON serialisation
type JTime struct {
	time.Time
}

func (t JTime) MarshalJSON() ([]byte, error) {

}
*/

/*
// what encoding.json does in Decode/Unmarshal
i, ok := val.(Marshaler)
if ok {
	data, err = i.Marshal()
} else {
	...
}
*/
