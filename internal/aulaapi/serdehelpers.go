package aulaapi

import (
	"encoding/json"
	"fmt"
)

// StringOrNumber is a custom type that deserializes from either a JSON string
// or number. The Aula API (.NET backend) is inconsistent about whether ID
// fields are serialized as strings or numbers.
type StringOrNumber string

// UnmarshalJSON implements json.Unmarshaler for StringOrNumber.
func (s *StringOrNumber) UnmarshalJSON(data []byte) error {
	// Try string first.
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*s = StringOrNumber(str)
		return nil
	}

	// Try number (int or float).
	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil {
		*s = StringOrNumber(num.String())
		return nil
	}

	// Try null.
	if string(data) == "null" {
		*s = ""
		return nil
	}

	return fmt.Errorf("StringOrNumber: cannot unmarshal %s", string(data))
}

// MarshalJSON implements json.Marshaler for StringOrNumber.
func (s StringOrNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// String returns the string value.
func (s StringOrNumber) String() string {
	return string(s)
}

// OptionalStringOrNumber is like StringOrNumber but handles null as nil.
type OptionalStringOrNumber struct {
	Value *string
}

// UnmarshalJSON implements json.Unmarshaler.
func (o *OptionalStringOrNumber) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.Value = nil
		return nil
	}

	// Try string.
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		o.Value = &str
		return nil
	}

	// Try number.
	var num json.Number
	if err := json.Unmarshal(data, &num); err == nil {
		s := num.String()
		o.Value = &s
		return nil
	}

	// Fallback: stringify whatever it is.
	s := string(data)
	o.Value = &s
	return nil
}

// MarshalJSON implements json.Marshaler.
func (o OptionalStringOrNumber) MarshalJSON() ([]byte, error) {
	if o.Value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*o.Value)
}
