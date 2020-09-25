package bertlv

import (
	"fmt"
	"reflect"
	"strings"
)

const structTagName = "bertlv"

// Raw is unparsed BER-TLV data
type Raw []byte

// Marshal converts a struct with bertlv tags to a BER-TLV encoded byte slice
func Marshal(v interface{}) ([]byte, error) {
	vType := reflect.TypeOf(v)
	for vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
	}
	if vType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("non-structs not supported, got %s", vType.Kind().String())
	}
	numFields := vType.NumField()
	for i := 0; i < numFields; i++ {
		field := vType.Field(i)
		// Only check tags on exported fields
		if string([]byte{field.Name[0]}) == strings.ToUpper(string([]byte{field.Name[0]})) {
			tag, exists := field.Tag.Lookup(structTagName)
			if !exists {
				return nil, fmt.Errorf("missing struct tag on field %s", field.Name)
			}
			_ = tag
		}
	}
	return nil, fmt.Errorf("not implemented")
}
