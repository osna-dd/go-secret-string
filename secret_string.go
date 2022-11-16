package secret

import "fmt"

type String struct {
	value string
}

func New(s string) String {
	return String{value: s}
}

func (s String) String() string {
	return "****"
}

func (s String) GoString() string {
	return "****"
}

func (s String) Get() string {
	return s.value
}

func (s String) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("marshal secret value")
}

type String2 string

func (s String2) String() string {
	return "****"
}

func (s String2) GoString() string {
	return "****"
}

func (s String2) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("marshal secret value")
}
