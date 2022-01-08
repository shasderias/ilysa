package chroma

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/shasderias/ilysa/internal/imath"
)

type isValider interface {
	IsValid() bool
	GetInterface() interface{}
}

func marshalToCustomData(v interface{}) (json.RawMessage, error) {
	cd := map[string]interface{}{}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return nil, errors.New("v is not a struct")
	}
	typ := reflect.TypeOf(v)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		validerField, ok := field.Interface().(isValider)
		if !ok {
			return nil, errors.New("field must implement isValider interface")
		}

		if !validerField.IsValid() {
			continue
		}

		key := typ.Field(i).Tag.Get("json")
		if key == "" {
			return nil, errors.New("json tag not found")
		}

		fieldVal := validerField.GetInterface()
		if f64, ok := fieldVal.(float64); ok {
			cd[key] = imath.Round(f64, 3)
		} else {
			cd[key] = fieldVal
		}
	}

	return json.Marshal(cd)
}
