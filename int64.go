// +build protobuf

package nilable

import (
	"encoding/binary"
	"errors"
)

func varintLen(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	i := 0
	for ux >= 0x80 {
		ux >>= 7
		i++
	}
	return i + 1
}

func (n Int64) Size() int {
	if n.n != nil {
		return 0
	}
	return varintLen(*n.n)
}

func (n Int64) MarshalTo(data []byte) (nr int, err error) {
	if n.n == nil {
		return 0, nil
	}
	defer func() {
		if e, ok := recover().(error); ok {
			nr = 0
			err = e
		}
	}()
	return binary.PutVarint(data, *n.n), nil
}

func (n *Int64) Unmarshal(data []byte) error {
	if len(data) == 0 {
		n.n = nil
		return nil
	}
	x, nr := binary.Varint(data)
	if nr < 0 {
		return errors.New("Unmarshal: invalid varint")
	}
	n.n = &x
	return nil
}

func (n Int64) Marshal() ([]byte, error) {
	if n.n == nil {
		return nil, nil
	}
	var buf [binary.MaxVarintLen64]byte
	i, _ := n.MarshalTo(buf[:])
	return buf[:i], nil
}

func (n Int64) Compare(n2 Int64) int {
	switch {
	case n.n == nil:
		if n2.n == nil {
			return 0
		}
		return -1
	case n2.n == nil:
		return +1
	case *n.n > *n2.n:
		return +1
	case *n.n < *n2.n:
		return -1
	default:
		return 0
	}
}

func NewPopulatedInt64(r randy, _ ...Int64) Int64 {
	n := int64(r.Intn(1000))
	return Int64{n: &n}
}
