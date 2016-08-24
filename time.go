// +build protobuf

package nilable

import "time"

func (t Time) Equal(t2 Time) bool {
	if t.t == nil {
		return t2.t == nil
	}
	return t2.t != nil && t.t.Equal(*t2.t)
}

func (t Time) Size() int {
	if t.t == nil {
		return 0
	}
	// https://golang.org/src/time/time.go?s=24443:24488#L832
	return 1 + 8 + 4 + 2
}

func (t Time) MarshalTo(data []byte) (int, error) {
	if t.t == nil {
		return 0, nil
	}
	b, err := t.t.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return copy(data, b), nil
}

func (t *Time) Unmarshal(data []byte) error {
	if t.t == nil {
		t.t = new(time.Time)
	}
	err := t.t.UnmarshalBinary(data)
	if err != nil {
		t.t = nil
		return err
	}
	return nil
}

func (t Time) Marshal() ([]byte, error) {
	if t.t == nil {
		return nil, nil
	}
	return t.t.MarshalBinary()
}

func (t Time) Compare(t2 Time) int {
	switch {
	case t.t == nil:
		if t2.t == nil {
			return 0
		}
		return -1
	case t2.t == nil:
		return +1
	case t.t.Before(*t2.t):
		return +1
	case t.t.Equal(*t2.t):
		return 0
	default:
		return -1
	}
}

func NewPopulatedTime(r randy, _ ...bool) *Time {
	const maxInt = int(^uint(0) >> 1)
	t := time.Unix(int64(r.Intn(maxInt)), int64(r.Intn(maxInt)))
	return &Time{t: &t}
}
