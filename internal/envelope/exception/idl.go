// Code generated by thriftrw v1.16.1. DO NOT EDIT.
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

package exception

import "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "exception",
	Package:  "go.uber.org/thriftrw/internal/envelope/exception",
	FilePath: "exception.thrift",
	SHA1:     "88105bcd404d4aee06542af9452f7cf76647ae98",
	Raw:      rawIDL,
}

const rawIDL = "enum ExceptionType {\n  UNKNOWN = 0\n  UNKNOWN_METHOD = 1\n  INVALID_MESSAGE_TYPE = 2\n  WRONG_METHOD_NAME = 3\n  BAD_SEQUENCE_ID = 4\n  MISSING_RESULT = 5\n  INTERNAL_ERROR = 6\n  PROTOCOL_ERROR = 7\n  INVALID_TRANSFORM = 8\n  INVALID_PROTOCOL = 9\n  UNSUPPORTED_CLIENT_TYPE = 10\n}\n\nexception TApplicationException {\n  1: optional string message\n  2: optional ExceptionType type\n}\n"
