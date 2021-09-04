package colorful

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// A HexColor is a Color that marshels to a hex triple of the form "#rrggbb".
// It implements the database/sql.Scanner, database/sql/driver.Value,
// encoding/json.Unmarshaler and encoding/json.Marshaler interfaces.
type HexColor Color

type errUnsupportedType struct {
	got  interface{}
	want reflect.Type
}

func (e errUnsupportedType) Error() string {
	return fmt.Sprintf("unsupported type: got %v, want a %s", e.got, e.want)
}

func (hc *HexColor) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errUnsupportedType{got: reflect.TypeOf(value), want: reflect.TypeOf("")}
	}
	c, err := ParseHex(s)
	if err != nil {
		return err
	}
	*hc = HexColor(c)
	return nil
}

func (hc *HexColor) Value() (driver.Value, error) {
	return Color(*hc).Hex(), nil
}

func (hc *HexColor) UnmarshalJSON(data []byte) error {
	var hexCode string
	if err := json.Unmarshal(data, &hexCode); err != nil {
		return err
	}

	var col, err = ParseHex(hexCode)
	if err != nil {
		return err
	}
	*hc = HexColor(col)
	return nil
}

func (hc HexColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(Color(hc).Hex())
}
