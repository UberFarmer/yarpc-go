// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package kv

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// KeyValue_SetValue_Args represents the arguments for the KeyValue.setValue function.
//
// The arguments for setValue are sent and received over the wire as this struct.
type KeyValue_SetValue_Args struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// ToWire translates a KeyValue_SetValue_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *KeyValue_SetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Key != nil {
		w, err = wire.NewValueString(*(v.Key)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Value != nil {
		w, err = wire.NewValueString(*(v.Value)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_SetValue_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_SetValue_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_SetValue_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_SetValue_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Key = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Value = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a KeyValue_SetValue_Args
// struct.
func (v *KeyValue_SetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	if v.Value != nil {
		fields[i] = fmt.Sprintf("Value: %v", *(v.Value))
		i++
	}

	return fmt.Sprintf("KeyValue_SetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_SetValue_Args match the
// provided KeyValue_SetValue_Args.
//
// This function performs a deep comparison.
func (v *KeyValue_SetValue_Args) Equals(rhs *KeyValue_SetValue_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Key, rhs.Key) {
		return false
	}
	if !_String_EqualsPtr(v.Value, rhs.Value) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_SetValue_Args.
func (v *KeyValue_SetValue_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Key != nil {
		enc.AddString("key", *v.Key)
	}
	if v.Value != nil {
		enc.AddString("value", *v.Value)
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *KeyValue_SetValue_Args) GetKey() (o string) {
	if v.Key != nil {
		return *v.Key
	}

	return
}

// IsSetKey returns true if Key is not nil.
func (v *KeyValue_SetValue_Args) IsSetKey() bool {
	return v.Key != nil
}

// GetValue returns the value of Value if it is set or its
// zero value if it is unset.
func (v *KeyValue_SetValue_Args) GetValue() (o string) {
	if v.Value != nil {
		return *v.Value
	}

	return
}

// IsSetValue returns true if Value is not nil.
func (v *KeyValue_SetValue_Args) IsSetValue() bool {
	return v.Value != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "setValue" for this struct.
func (v *KeyValue_SetValue_Args) MethodName() string {
	return "setValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *KeyValue_SetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// KeyValue_SetValue_Helper provides functions that aid in handling the
// parameters and return values of the KeyValue.setValue
// function.
var KeyValue_SetValue_Helper = struct {
	// Args accepts the parameters of setValue in-order and returns
	// the arguments struct for the function.
	Args func(
		key *string,
		value *string,
	) *KeyValue_SetValue_Args

	// IsException returns true if the given error can be thrown
	// by setValue.
	//
	// An error can be thrown by setValue only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for setValue
	// given the error returned by it. The provided error may
	// be nil if setValue did not fail.
	//
	// This allows mapping errors returned by setValue into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// setValue
	//
	//   err := setValue(args)
	//   result, err := KeyValue_SetValue_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from setValue: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*KeyValue_SetValue_Result, error)

	// UnwrapResponse takes the result struct for setValue
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if setValue threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := KeyValue_SetValue_Helper.UnwrapResponse(result)
	UnwrapResponse func(*KeyValue_SetValue_Result) error
}{}

func init() {
	KeyValue_SetValue_Helper.Args = func(
		key *string,
		value *string,
	) *KeyValue_SetValue_Args {
		return &KeyValue_SetValue_Args{
			Key:   key,
			Value: value,
		}
	}

	KeyValue_SetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	KeyValue_SetValue_Helper.WrapResponse = func(err error) (*KeyValue_SetValue_Result, error) {
		if err == nil {
			return &KeyValue_SetValue_Result{}, nil
		}

		return nil, err
	}
	KeyValue_SetValue_Helper.UnwrapResponse = func(result *KeyValue_SetValue_Result) (err error) {
		return
	}

}

// KeyValue_SetValue_Result represents the result of a KeyValue.setValue function call.
//
// The result of a setValue execution is sent and received over the wire as this struct.
type KeyValue_SetValue_Result struct {
}

// ToWire translates a KeyValue_SetValue_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *KeyValue_SetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_SetValue_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_SetValue_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_SetValue_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_SetValue_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a KeyValue_SetValue_Result
// struct.
func (v *KeyValue_SetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("KeyValue_SetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_SetValue_Result match the
// provided KeyValue_SetValue_Result.
//
// This function performs a deep comparison.
func (v *KeyValue_SetValue_Result) Equals(rhs *KeyValue_SetValue_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_SetValue_Result.
func (v *KeyValue_SetValue_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "setValue" for this struct.
func (v *KeyValue_SetValue_Result) MethodName() string {
	return "setValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *KeyValue_SetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
