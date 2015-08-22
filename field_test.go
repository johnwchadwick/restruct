package restruct

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var intType = reflect.TypeOf(int(0))
var boolType = reflect.TypeOf(false)
var strType = reflect.TypeOf(string(""))

func TestFieldsFromStruct(t *testing.T) {
	tests := []struct {
		input  interface{}
		fields []Field
	}{
		{
			struct {
				Simple int
			}{},
			[]Field{
				Field{"Simple", 0, true, intType, intType, nil, true},
			},
		},
		{
			struct {
				Before int
				During string `struct:"-"`
				After  bool
			}{},
			[]Field{
				Field{"Before", 0, true, intType, intType, nil, true},
				Field{"After", 2, true, boolType, boolType, nil, true},
			},
		},
		{
			struct {
				FixedStr string `struct:"[64]byte"`
				LSBInt   int    `struct:"uint32,little"`
			}{},
			[]Field{
				Field{"FixedStr", 0, true, reflect.TypeOf([64]byte{}), strType, nil, true},
				Field{"LSBInt", 1, true, reflect.TypeOf(uint32(0)), intType, binary.LittleEndian, true},
			},
		},
	}

	for _, test := range tests {
		fields := FieldsFromStruct(reflect.TypeOf(test.input))
		assert.Equal(t, fields, test.fields)
	}
}

func TestFieldsFromNonStructPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Non-struct did not panic.")
		}
	}()
	FieldsFromStruct(reflect.TypeOf(0))
}

func TestIsTypeTrivial(t *testing.T) {
	tests := []struct {
		input   interface{}
		trivial bool
	}{
		{int8(0), true},
		{int16(0), true},
		{int32(0), true},
		{int64(0), true},
		{[0]int8{}, true},
		{[]int8{}, false},
		{struct{}{}, true},
		{struct{ int8 }{}, true},
		{struct{ a []int8 }{[]int8{}}, false},
		{struct{ a [0]int8 }{[0]int8{}}, true},
		{(*interface{})(nil), false},
	}

	for _, test := range tests {
		assert.Equal(t, test.trivial, IsTypeTrivial(reflect.TypeOf(test.input)))
	}
}
