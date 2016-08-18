// +build protobuf

package nilable

func (s String) Size() int {
	if s.s == nil {
		return 0
	}
	return len(*s.s)
}

func (s String) MarshalTo(data []byte) (int, error) {
	if s.s == nil {
		return 0, nil
	}
	return copy(data, *s.s), nil
}

func (s *String) Unmarshal(data []byte) error {
	if data == nil {
		s.s = nil
	} else {
		str := string(data)
		s.s = &str
	}
	return nil
}

func (s String) Marshal() ([]byte, error) {
	if s.s == nil {
		return nil, nil
	}
	b := make([]byte, len(*s.s))
	_, err := s.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s String) Compare(s2 String) int {
	switch {
	case s.s == nil:
		if s2.s == nil {
			return 0
		}
		return -1
	case s2.s == nil:
		return +1
	case *s.s > *s2.s:
		return +1
	case *s.s < *s2.s:
		return -1
	default:
		return 0
	}
}

type randy interface {
	Intn(n int) int
}

func NewPopulatedString(r randy, _ ...bool) String {
	buf := make([]byte, r.Intn(50))
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(r.Intn(255))
	}
	str := string(buf)
	return String{s: &str}
}
