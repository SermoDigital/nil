package nilable

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type Nilable interface {
	IsNil() bool
}

type String struct {
	s *string
}

func NewString(s string) String {
	return String{s: &s}
}

func (n String) IsNil() bool {
	return n.s == nil
}

func (n String) String() string {
	if n.IsNil() {
		return ""
	}
	return *n.s
}

func (n *String) Scan(value interface{}) error {
	if value == nil {
		return errors.New("nilable.String.Scan: value is nil")
	}
	switch v := value.(type) {
	case string:
		n.s = &v
	case []byte:
		vb := string(v)
		n.s = &vb
	default:
		return errors.New("nilable.String.Scan: value is an incorrect type")
	}
	return nil
}

func (n String) Value() (driver.Value, error) {
	if n.IsNil() {
		return nil, nil
	}
	return n.String(), nil
}

func (n String) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.s)
}

type Bool struct {
	b *bool
}

func NewBool(b bool) Bool {
	return Bool{b: &b}
}

func (n Bool) IsNil() bool {
	return n.b == nil
}

func (n Bool) Bool() bool {
	if n.IsNil() {
		return false
	}
	return *n.b
}

func (n *Bool) Scan(value interface{}) error {
	if value == nil {
		return errors.New("nilable.Bool.Scan: value is nil")
	}
	switch v := value.(type) {
	case bool:
		n.b = &v
	default:
		return errors.New("nilable.Bool.Scan: value is an incorrect type")
	}
	return nil
}

func (n Bool) Value() (driver.Value, error) {
	if n.IsNil() {
		return nil, nil
	}
	return n.Bool(), nil
}

func (n Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Bool())
}

func (n *Bool) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.b)
}

func (n Bool) String() string {
	switch {
	case n.IsNil():
		return "<nil>"
	case n.Bool():
		return "true"
	default:
		return "false"
	}
}

type Time struct {
	t *time.Time
}

func NewTime(t time.Time) Time {
	return Time{t: &t}
}

func (n Time) IsNil() bool {
	return n.t == nil
}

func (n Time) Time() time.Time {
	if n.IsNil() {
		return time.Time{}
	}
	return time.Time(*n.t)
}

func (n *Time) Scan(value interface{}) error {
	if value == nil {
		return errors.New("nilable.Time.Scan: value is nil")
	}
	switch v := value.(type) {
	case time.Time:
		n.t = &v
	default:
		return errors.New("nilable.Time.Scan: value is an incorrect type")
	}
	return nil
}

func (n Time) String() string {
	if n.IsNil() {
		return "<nil>"
	}
	return n.t.String()
}

func (n Time) Value() (driver.Value, error) {
	if n.IsNil() {
		return nil, nil
	}
	return n.Time(), nil
}

func (n Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Time())
}

func (n *Time) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.t)
}

type Int64 struct {
	n *int64
}

func NewInt64(n int64) Int64 {
	return Int64{n: &n}
}

func (n Int64) IsNil() bool {
	return n.n == nil
}

func (n Int64) Int64() int64 {
	if n.IsNil() {
		return 0
	}
	return *n.n
}

func (n *Int64) Scan(value interface{}) error {
	if value == nil {
		return errors.New("nilable.Int64.Scan: value is nil")
	}
	switch v := value.(type) {
	case int64:
		n.n = &v
	case int:
		nv := int64(v)
		n.n = &nv
	case int32:
		nv := int64(v)
		n.n = &nv
	case int16:
		nv := int64(v)
		n.n = &nv
	case int8:
		nv := int64(v)
		n.n = &nv
	default:
		return errors.New("nilable.Int64.Scan: value is an incorrect type")
	}
	return nil
}

func (n Int64) Value() (driver.Value, error) {
	if n.IsNil() {
		return nil, nil
	}
	return n.Int64(), nil
}

func (n Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Int64())
}

func (n *Int64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.n)
}

func (n Int64) String() string {
	if n.IsNil() {
		return "<nil>"
	}
	return strconv.FormatInt(*n.n, 10)
}

var (
	_ Nilable = (*String)(nil)
	_ Nilable = (*Bool)(nil)
	_ Nilable = (*Bool)(nil)
	_ Nilable = (*Int64)(nil)
	_ Nilable = (*Time)(nil)
)
