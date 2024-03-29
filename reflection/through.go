package reflection

import "reflect"

func Through(x interface{}, fn func(entry string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
