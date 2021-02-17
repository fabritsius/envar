package envar

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Fill uses environment variables or defaults to assign values to element of a struct.
func Fill(vars interface{}) error {
	vals := reflect.ValueOf(vars).Elem()
	typs := reflect.TypeOf(vars).Elem()

	for i := 0; i < vals.NumField(); i++ {
		fieldType := typs.Field(i)
		fieldVal := vals.FieldByName(fieldType.Name)

		strVal, ok := searchEnv(fieldType)
		if !ok {
			return fmt.Errorf("envar: nothing to assign to '%s' field", fieldType.Name)
		}

		switch fieldVal.Kind() {
		case reflect.String:
			fieldVal.SetString(strVal)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			numVal, err := strconv.ParseInt(strVal, 10, 64)
			if err != nil {
				return conversionError(fieldType.Name, strVal, "int64")
			}
			fieldVal.SetInt(numVal)
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(strVal)
			fmt.Println(strVal, boolVal)
			if err != nil {
				return conversionError(fieldType.Name, strVal, "bool")
			}
			fieldVal.SetBool(boolVal)
		default:
			return fmt.Errorf("envar: field '%s' should be of type 'string', 'int' or 'bool'", fieldType.Name)
		}
	}
	return nil
}

func searchEnv(fieldType reflect.StructField) (string, bool) {
	// check if there is an environment variable for this field
	if envKey, ok := fieldType.Tag.Lookup("env"); ok {
		if envVal, ok := os.LookupEnv(envKey); ok {
			return envVal, true
		}
	}
	// try to get the default value
	if defaultVal, ok := fieldType.Tag.Lookup("default"); ok {
		return defaultVal, true
	}
	return "", false
}

func conversionError(fieldName, strVal, desired string) error {
	return fmt.Errorf("envar: field '%s': cannot convert '%s' to type '%s'", fieldName, strVal, desired)
}
