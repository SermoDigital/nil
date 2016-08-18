package nilable

import "errors"

func (b Bool) Size() int {
	if b.b == nil {
		return 0
	}
	return 1
}

func (b Bool) MarshalTo(data []byte) (int, error) {
	if b.b == nil {
		return 0, nil
	}
	if *b.b {
		return copy(data, []byte{1}), nil
	}
	return copy(data, []byte{0}), nil
}

func (b *Bool) Unmarshal(data []byte) error {
	switch len(data) {
	case 0:
		b.b = nil
	case 1:
		t := data[0] == 1
		b.b = &t
	default:
		return errors.New("Unmarshal: invalid data length")
	}
	return nil
}

func (b Bool) Marshal() ([]byte, error) {
	var buf [1]byte
	_, err := b.MarshalTo(buf[:])
	if err != nil {
		return nil, err
	}
	return buf[:], nil
}

func (b Bool) Compare(b2 Bool) int {
	switch {
	case b.b == nil:
		if b2.b == nil {
			return 0
		}
		return -1
	case b2.b == nil:
		return +1
	default:
		if *b.b != *b2.b {
			if *b.b {
				return +1
			}
			return -1
		}
		return 0
	}
}

func NewPopulatedBool(r randy, _ ...bool) Bool {
	b := r.Intn(1000)%2 == 0
	return Bool{b: &b}
}
