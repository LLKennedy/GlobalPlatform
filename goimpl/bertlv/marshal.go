package bertlv

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

const structTagName = "bertlv"

// Raw is unparsed BER-TLV data
type Raw []byte

// Marshal converts a struct with bertlv tags to a BER-TLV encoded byte slice
func Marshal(v interface{}) (encoded []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", err)
		}
	}()
	vType := reflect.TypeOf(v)
	vVal := reflect.ValueOf(v)
	for vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
		vVal = vVal.Elem()
	}
	if vType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("non-structs not supported, got %s", vType.Kind().String())
	}
	buf := bytes.NewBuffer(nil)
	numFields := vType.NumField()
	for i := 0; i < numFields; i++ {
		field := vType.Field(i)
		fieldVal := vVal.FieldByName(field.Name)
		// Only check tags on exported fields
		if string([]byte{field.Name[0]}) == strings.ToUpper(string([]byte{field.Name[0]})) {
			tag, exists := field.Tag.Lookup(structTagName)
			if exists {
				parsedTag, omitEmpty := parseStructTag(tag, field.Name)
				obj := Object{
					Tag: parsedTag,
				}
				isNil := false
				if field.Type.Kind() == reflect.Ptr {

				}
				switch field.Type.Kind() {
				case reflect.Ptr:
				case reflect.Struct:
				case reflect.Bool:
				case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
				case reflect.Slice:
				}

			}
		}
	}
	return nil, fmt.Errorf("not implemented")
}

func parseStructTag(in, name string) (tag Tag, omitEmpty bool) {
	elems := strings.Split(in, ",")
	switch len(elems) {
	case 2:
		if elems[1] != "omitempty" {
			panic(fmt.Sprintf("second part of tag must be omitempty if it is present, found %s", elems[1]))
		}
		omitEmpty = true
		fallthrough
	case 1:
		tagBytes, err := hex.DecodeString(elems[0])
		if err != nil {
			panic(fmt.Sprintf("invalid tag hex in field %s: %v", name, err))
		}
		tagUint := binary.BigEndian.Uint64(tagBytes)
		tag = TagFromUintForced(tagUint)
		return
	default:
		panic(fmt.Sprintf("invalid struct tag on field %s: must have 1 or 2 comma-separated elements, found %d", name, len(elems)))
	}
}
