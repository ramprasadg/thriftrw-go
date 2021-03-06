// Code generated by thriftrw v1.17.0. DO NOT EDIT.
// @generated

package enums

import thriftreflect "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "enums",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/enums",
	FilePath: "enums.thrift",
	SHA1:     "16921e977d77e214a9bac398d279f60889b055de",
	Raw:      rawIDL,
}

const rawIDL = "enum EmptyEnum {}\n\nenum EnumDefault {\n    Foo, Bar, Baz\n}\n\nenum EnumWithValues {\n    X = 123,\n    Y = 456,\n    Z = 789,\n}\n\nenum EnumWithDuplicateValues {\n    P, // 0\n    Q = -1,\n    R, // 0\n}\n\n// enum with item names conflicting with those of another enum\nenum EnumWithDuplicateName {\n    A, B, C, P, Q, R, X, Y, Z\n}\n\n// Enum treated as optional inside a struct\nstruct StructWithOptionalEnum {\n    1: optional EnumDefault e\n}\n\n/**\n * Kinds of records stored in the database.\n */\nenum RecordType {\n  /** Name of the user. */\n  NAME,\n\n  /**\n   * Home address of the user.\n   *\n   * This record is always present.\n   */\n  HOME_ADDRESS,\n\n  /**\n   * Home address of the user.\n   *\n   * This record may not be present.\n   */\n  WORK_ADDRESS\n}\n\nenum lowerCaseEnum {\n    containing, lower_case, items\n}\n\n// EnumWithLabel use label name in serialization/deserialization\nenum EnumWithLabel {\n    USERNAME (go.label = \"surname\"),\n    PASSWORD (go.label = \"hashed_password\"),\n    SALT (go.label = \"\"),\n    SUGAR (go.label),\n    relay (go.label = \"RELAY\")\n    NAIVE4_N1 (go.label = \"function\")\n\n}\n\n// collision with RecordType_Values() function.\nenum RecordType_Values { FOO, BAR }\n"
