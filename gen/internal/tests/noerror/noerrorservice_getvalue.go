// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

package noerror

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// NoErrorService_GetValue_Args represents the arguments for the NoErrorService.getValue function.
//
// The arguments for getValue are sent and received over the wire as this struct.
type NoErrorService_GetValue_Args struct {
	Key *Key `json:"key,omitempty"`
}

// ToWire translates a NoErrorService_GetValue_Args struct into a Thrift-level intermediate
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
func (v *NoErrorService_GetValue_Args) ToWire() (wire.Value, error) {
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

// FromWire deserializes a NoErrorService_GetValue_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a NoErrorService_GetValue_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v NoErrorService_GetValue_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *NoErrorService_GetValue_Args) FromWire(w wire.Value) error {
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

// String returns a readable string representation of a NoErrorService_GetValue_Args
// struct.
func (v *NoErrorService_GetValue_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}

	return fmt.Sprintf("NoErrorService_GetValue_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this NoErrorService_GetValue_Args match the
// provided NoErrorService_GetValue_Args.
//
// This function performs a deep comparison.
func (v *NoErrorService_GetValue_Args) Equals(rhs *NoErrorService_GetValue_Args) bool {
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
// fast logging of NoErrorService_GetValue_Args.
func (v *NoErrorService_GetValue_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
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
func (v *NoErrorService_GetValue_Args) GetKey() (o Key) {
	if v != nil && v.Key != nil {
		return *v.Key
	}

	return
}

// IsSetKey returns true if Key is not nil.
func (v *NoErrorService_GetValue_Args) IsSetKey() bool {
	return v != nil && v.Key != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "getValue" for this struct.
func (v *NoErrorService_GetValue_Args) MethodName() string {
	return "getValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *NoErrorService_GetValue_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// NoErrorService_GetValue_Helper provides functions that aid in handling the
// parameters and return values of the NoErrorService.getValue
// function.
var NoErrorService_GetValue_Helper = struct {
	// Args accepts the parameters of getValue in-order and returns
	// the arguments struct for the function.
	Args func(
		key *Key,
	) *NoErrorService_GetValue_Args

	// IsException returns true if the given value can be thrown
	// by getValue.
	//
	// An exception can be thrown by getValue only if the
	// type was mentioned in the 'throws' section of the IDL.
	IsException func(interface{}) bool

	// WrapResponse returns the result struct for getValue
	// given its return value.
	//
	// This allows mapping values and exceptions returned by
	// getValue into a serializable result struct.
	// WrapResponse returns an error if the provided
	// value cannot be returned by getValue.
	//
	//   value, err := getValue(args)
	//   result, err := NoErrorService_GetValue_Helper.WrapResponse(value)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from getValue: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(interface{}) (*NoErrorService_GetValue_Result, error)

	// UnwrapResponse takes the result struct for getValue
	// and returns the value or exception returned by it.
	//
	// The error is non-nil only if the result is unrecognized.
	//
	//   result := deserialize(bytes)
	//   value, exception, err := NoErrorService_GetValue_Helper.UnwrapResponse(result)
	UnwrapResponse func(*NoErrorService_GetValue_Result) (Key, interface{}, error)
}{}

func init() {
	NoErrorService_GetValue_Helper.Args = func(
		key *Key,
	) *NoErrorService_GetValue_Args {
		return &NoErrorService_GetValue_Args{
			Key: key,
		}
	}

	NoErrorService_GetValue_Helper.IsException = func(err interface{}) bool {
		switch err.(type) {
		case *NoErrorException, NoErrorException:
			return true
		default:
			return false
		}
	}

	NoErrorService_GetValue_Helper.WrapResponse = func(val interface{}) (*NoErrorService_GetValue_Result, error) {
		if retVal, ok := val.(Key); ok {
			return &NoErrorService_GetValue_Result{Success: &retVal}, nil
		}

		switch e := val.(type) {
		case *NoErrorException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for NoErrorService_GetValue_Result.DoesNotExist")
			}
			return &NoErrorService_GetValue_Result{DoesNotExist: e}, nil
		}

		return nil, fmt.Errorf("WrapResponse received an unrecognized type for NoErrorService_GetValue_Result: %v", val)
	}
	NoErrorService_GetValue_Helper.UnwrapResponse = func(result *NoErrorService_GetValue_Result) (success Key, exception interface{}, err error) {
		if result.DoesNotExist != nil {
			exception = result.DoesNotExist
			return
		}

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// NoErrorService_GetValue_Result represents the result of a NoErrorService.getValue function call.
//
// The result of a getValue execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type NoErrorService_GetValue_Result struct {
	// Value returned by getValue after a successful execution.
	Success      *Key              `json:"success,omitempty"`
	DoesNotExist *NoErrorException `json:"doesNotExist,omitempty"`
}

// ToWire translates a NoErrorService_GetValue_Result struct into a Thrift-level intermediate
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
func (v *NoErrorService_GetValue_Result) ToWire() (wire.Value, error) {
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
		return wire.Value{}, fmt.Errorf("NoErrorService_GetValue_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a NoErrorService_GetValue_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a NoErrorService_GetValue_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v NoErrorService_GetValue_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *NoErrorService_GetValue_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x Key
				x, err = _Key_Read(field.Value)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _NoErrorException_Read(field.Value)
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
		return fmt.Errorf("NoErrorService_GetValue_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a NoErrorService_GetValue_Result
// struct.
func (v *NoErrorService_GetValue_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}

	return fmt.Sprintf("NoErrorService_GetValue_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this NoErrorService_GetValue_Result match the
// provided NoErrorService_GetValue_Result.
//
// This function performs a deep comparison.
func (v *NoErrorService_GetValue_Result) Equals(rhs *NoErrorService_GetValue_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_Key_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of NoErrorService_GetValue_Result.
func (v *NoErrorService_GetValue_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddString("success", (string)(*v.Success))
	}
	if v.DoesNotExist != nil {
		err = multierr.Append(err, enc.AddObject("doesNotExist", v.DoesNotExist))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *NoErrorService_GetValue_Result) GetSuccess() (o Key) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *NoErrorService_GetValue_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetDoesNotExist returns the value of DoesNotExist if it is set or its
// zero value if it is unset.
func (v *NoErrorService_GetValue_Result) GetDoesNotExist() (o *NoErrorException) {
	if v != nil && v.DoesNotExist != nil {
		return v.DoesNotExist
	}

	return
}

// IsSetDoesNotExist returns true if DoesNotExist is not nil.
func (v *NoErrorService_GetValue_Result) IsSetDoesNotExist() bool {
	return v != nil && v.DoesNotExist != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "getValue" for this struct.
func (v *NoErrorService_GetValue_Result) MethodName() string {
	return "getValue"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *NoErrorService_GetValue_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
