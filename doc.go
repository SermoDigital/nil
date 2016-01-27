// Package nilable is an easy way to easily use Go types while still
// distinguishing between an absent value and a zero-type.
//
// This distinction (absent versus zero) can occur in APIs that unmarshal
// a JSON payload into a specific type. If the payload does not set specific
// JSON fields, the Go type it's marshaled into will set the unspecified fields
// to their zero values.
//
// This is remedied in one of two ways:
//
// First, the API could use "replace" semantics instead of "edit." This means
// the object or data sent through the API is replacing what it represents on
// the server. In practical terms, this could be represented by first getting
// the entire object from the server, modifying the necessary fields, and then
// sending the entire object back.
//
// While this works, it requires two round-trips to the server and could allow
// for race conditions, depending how the API is structured.
//
// The second remedy is to use something similar to this package: nilable values.
// Nilable values would solve the previous problem by creating a distinction
// between absent and zero. This means instead of two round-trips to the server,
// the client would only have to send one request containing the modified fields.
//
// All basic types in package nilable implement the Nilable interface.
//
// The basic structure of a nilable type is a struct containing a pointer to the
// native type.
//
// For example:
//
// 		type String struct {
// 			s *string
// 		}
//
// Each type has a method which returns the underlying type. By convention,
// this method is the name of the Nilable type, which is also the exported
// version of the name of the underlying type.
//
// 		int64 -> Int64
// 		string -> String
// 		bool -> Bool
//
// Each type has a String method which returns the canonical string version
// of the type. (As it would normally be printed using the "%s" verb.)
// One caveat is nil -- if the type is nil, "<nil>" will be returned.
//
// Each type should be passed by value like a normal Go type. (I.e., do not
// pass around *T.)
package nilable
