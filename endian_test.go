// Copyright (c) 2014 Naoya SATO
//
// Originally code is:
// Copyright (c) 2009 The Go Authors.
// Released under the BSD License.

package endian

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
	"testing"
)

type Struct struct {
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
	Array      [4]uint8
}

var s = Struct{
	0x01,
	0x0203,
	0x04050607,
	0x08090a0b0c0d0e0f,
	0x10,
	0x1112,
	0x13141516,
	0x1718191a1b1c1d1e,

	math.Float32frombits(0x1f202122),
	math.Float64frombits(0x232425262728292a),
	complex(
		math.Float32frombits(0x2b2c2d2e),
		math.Float32frombits(0x2f303132),
	),
	complex(
		math.Float64frombits(0x333435363738393a),
		math.Float64frombits(0x3b3c3d3e3f404142),
	),

	[4]uint8{0x43, 0x44, 0x45, 0x46},
}

var middle = []byte{
	1,
	3, 2,
	5, 4, 7, 6,
	9, 8, 11, 10, 13, 12, 15, 14,
	16,
	18, 17,
	20, 19, 22, 21,
	24, 23, 26, 25, 28, 27, 30, 29,

	32, 31, 34, 33,
	36, 35, 38, 37, 40, 39, 42, 41,
	44, 43, 46, 45, 48, 47, 50, 49,
	52, 51, 54, 53, 56, 55, 58, 57, 60, 59, 62, 61, 64, 63, 66, 65,

	67, 68, 69, 70,
}

func checkResult(t *testing.T, dir string, order binary.ByteOrder, err error, have, want interface{}) {
	if err != nil {
		t.Errorf("%v %v: %v", dir, order, err)
		return
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("%v %v:\n\thave %+v\n\twant %+v", dir, order, have, want)
	}
}

func testRead(t *testing.T, order binary.ByteOrder, b []byte, s1 interface{}) {
	var s2 Struct
	err := binary.Read(bytes.NewReader(b), order, &s2)
	checkResult(t, "Read", order, err, s2, s1)
}

func testWrite(t *testing.T, order binary.ByteOrder, b []byte, s1 interface{}) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, s1)
	checkResult(t, "Write", order, err, buf.Bytes(), b)
}

func TestMiddleEndianRead(t *testing.T)     { testRead(t, MiddleEndian, middle, s) }
func TestMiddleEndianWrite(t *testing.T)    { testWrite(t, MiddleEndian, middle, s) }
func TestMiddleEndianPtrWrite(t *testing.T) { testWrite(t, MiddleEndian, middle, &s) }
