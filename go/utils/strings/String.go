package strings

import "bytes"

// String is a wrapper over buff.bytes to make it seamless to concatenate strings
type String struct {
	buff *bytes.Buffer
}

// New construct a new String instance and initialize the buff with the input string
func New(s string) *String {
	_String := &String{}
	_String.init()
	_String.buff.WriteString(s)
	return _String
}

// init initialize the buff if needed
func (s *String) init() {
	if s.buff == nil {
		s.buff = &bytes.Buffer{}
	}
}

// Add concatenate a string to this String instance
func (s *String) Add(str string) *String {
	s.init()
	s.buff.WriteString(str)
	return s
}

// Join concatenate a String instance to this String instance
func (s *String) Join(other *String) *String {
	s.init()
	s.buff.Write(other.buff.Bytes())
	return s
}

// String convert the String instance buffer to primitive string
func (s *String) String() string {
	s.init()
	return s.buff.String()
}

// IsBlank return is this String instance is blank
func (s *String) IsBlank() bool {
	s.init()
	return s.buff.Len() == 0
}
