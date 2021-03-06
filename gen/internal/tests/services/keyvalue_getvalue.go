// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

package services

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	exceptions "go.uber.org/thriftrw/gen/internal/tests/exceptions"
	unions "go.uber.org/thriftrw/gen/internal/tests/unions"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// KeyValue_GetValue_Args represents the arguments for the KeyValue.getValue function.
//
// The arguments for getValue are sent and received over the wire as this struct.
type KeyValue_GetValue_Args struct {
	Key *Key `json:"key,omitempty"`
}

// ToWire translates a KeyValue_GetValue_Args struct into a Thrift-level intermediate
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
func (v *KeyValue_GetValue_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Key != nil {
		w, err = v.Key.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_GetValue_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_GetValue_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_GetValue_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_GetValue_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x Key
				x, err = _Key_Read(field.Value)
				v.Key = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a KeyValue_GetValue_Args
// struct.
func (v *KeyValue_GetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}

	return fmt.Sprintf("KeyValue_GetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_GetValue_Args match the
// provided KeyValue_GetValue_Args.
//
// This function performs a deep comparison.
func (v *KeyValue_GetValue_Args) Equals(rhs *KeyValue_GetValue_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_Key_EqualsPtr(v.Key, rhs.Key) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_GetValue_Args.
func (v *KeyValue_GetValue_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Key != nil {
		enc.AddString("key", (string)(*v.Key))
	}
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetValue_Args) GetKey() (o Key) {
	if v != nil && v.Key != nil {
		return *v.Key
	}

	return
}

// IsSetKey returns true if Key is not nil.
func (v *KeyValue_GetValue_Args) IsSetKey() bool {
	return v != nil && v.Key != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "getValue" for this struct.
func (v *KeyValue_GetValue_Args) MethodName() string {
	return "getValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *KeyValue_GetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// KeyValue_GetValue_Helper provides functions that aid in handling the
// parameters and return values of the KeyValue.getValue
// function.
var KeyValue_GetValue_Helper = struct {
	// Args accepts the parameters of getValue in-order and returns
	// the arguments struct for the function.
	Args func(
		key *Key,
	) *KeyValue_GetValue_Args

	// IsException returns true if the given error can be thrown
	// by getValue.
	//
	// An error can be thrown by getValue only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for getValue
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// getValue into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by getValue
	//
	//   value, err := getValue(args)
	//   result, err := KeyValue_GetValue_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from getValue: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*unions.ArbitraryValue, error) (*KeyValue_GetValue_Result, error)

	// UnwrapResponse takes the result struct for getValue
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if getValue threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := KeyValue_GetValue_Helper.UnwrapResponse(result)
	UnwrapResponse func(*KeyValue_GetValue_Result) (*unions.ArbitraryValue, error)
}{}

func init() {
	KeyValue_GetValue_Helper.Args = func(
		key *Key,
	) *KeyValue_GetValue_Args {
		return &KeyValue_GetValue_Args{
			Key: key,
		}
	}

	KeyValue_GetValue_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		default:
			return false
		}
	}

	KeyValue_GetValue_Helper.WrapResponse = func(success *unions.ArbitraryValue, err error) (*KeyValue_GetValue_Result, error) {
		if err == nil {
			return &KeyValue_GetValue_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_GetValue_Result.DoesNotExist")
			}
			return &KeyValue_GetValue_Result{DoesNotExist: e}, nil
		}

		return nil, err
	}
	KeyValue_GetValue_Helper.UnwrapResponse = func(result *KeyValue_GetValue_Result) (success *unions.ArbitraryValue, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// KeyValue_GetValue_Result represents the result of a KeyValue.getValue function call.
//
// The result of a getValue execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type KeyValue_GetValue_Result struct {
	// Value returned by getValue after a successful execution.
	Success      *unions.ArbitraryValue            `json:"success,omitempty"`
	DoesNotExist *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
}

// ToWire translates a KeyValue_GetValue_Result struct into a Thrift-level intermediate
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
func (v *KeyValue_GetValue_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a KeyValue_GetValue_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_GetValue_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_GetValue_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_GetValue_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _ArbitraryValue_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("KeyValue_GetValue_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a KeyValue_GetValue_Result
// struct.
func (v *KeyValue_GetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}

	return fmt.Sprintf("KeyValue_GetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this KeyValue_GetValue_Result match the
// provided KeyValue_GetValue_Result.
//
// This function performs a deep comparison.
func (v *KeyValue_GetValue_Result) Equals(rhs *KeyValue_GetValue_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_GetValue_Result.
func (v *KeyValue_GetValue_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.DoesNotExist != nil {
		err = multierr.Append(err, enc.AddObject("doesNotExist", v.DoesNotExist))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetValue_Result) GetSuccess() (o *unions.ArbitraryValue) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *KeyValue_GetValue_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetDoesNotExist returns the value of DoesNotExist if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetValue_Result) GetDoesNotExist() (o *exceptions.DoesNotExistException) {
	if v != nil && v.DoesNotExist != nil {
		return v.DoesNotExist
	}

	return
}

// IsSetDoesNotExist returns true if DoesNotExist is not nil.
func (v *KeyValue_GetValue_Result) IsSetDoesNotExist() bool {
	return v != nil && v.DoesNotExist != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "getValue" for this struct.
func (v *KeyValue_GetValue_Result) MethodName() string {
	return "getValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *KeyValue_GetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
