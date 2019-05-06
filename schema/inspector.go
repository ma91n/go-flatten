package schema

import (
	"fmt"
	"reflect"
	"errors"
)

type Inspector struct {
}

func (i *Inspector) Lookup(m map[string]interface{}) ([]string, error) {

	keys := make([]string, 0)

	for k, v := range m {
		switch tv := v.(type) {
		case []interface{}:
			for _, item := range tv {
				mv, ok := item.(map[string]interface{})
				if !ok {
					return nil, errors.New(fmt.Sprintf("inspection error %v", v))
				}

				nestKeys, err := i.Lookup(mv)
				if err != nil {
					return nil, err
				}
				keys = append(keys, nestKeys...)
			}
		case map[string]interface{}:
			nestKeys, err := i.Lookup(tv)
			if err != nil {
				return nil, err
			}
			keys = append(keys, nestKeys...)

		default:
			switch reflect.TypeOf(k).Kind() {
			case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
				fmt.Printf("TypeSwitch was unable to identify this item. Reflect says %#v is a slice, array, map, or channel.\n", k)
			default:
				fmt.Printf("Type handler not implemented: %#v\n", k)
			}
		case string:
			keys = append(keys, k)
		}
	}
	return keys, nil
}
