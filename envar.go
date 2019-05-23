package envar

import (
	"fmt"
	"os"
	"reflect"
)

// Fill uses environment variables or defaults to assign values to element of a struct.
func Fill(vars interface{}) error {
	vals := reflect.ValueOf(vars).Elem()
	typs := reflect.TypeOf(vars).Elem()

	for i := 0; i < vals.NumField(); i++ {
		fieldType := typs.Field(i)
		fieldVal := vals.FieldByName(fieldType.Name)
		// check if field is a string
		if fieldVal.Kind() != reflect.String {
			return fmt.Errorf("envar: field '%s' should be of type 'string'", fieldType.Name)
		}
		// check if there is an environment variable for this field
		if envKey, ok := fieldType.Tag.Lookup("env"); ok {
			if envVal, ok := os.LookupEnv(envKey); ok {
				fieldVal.SetString(envVal)
				continue
			}
		}
		// try to get the default value
		if defaultVal, ok := fieldType.Tag.Lookup("default"); ok {
			fieldVal.SetString(defaultVal)
		} else {
			return fmt.Errorf("envar: nothing to assign to '%s' field", fieldType.Name)
		}
	}
	return nil
}
