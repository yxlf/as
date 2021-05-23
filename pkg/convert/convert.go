package convert

import (
	"strconv"
)

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (v int, err error) {
	v, err = strconv.Atoi(s.String())
	return
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) Uint32() (v uint32, err error) {
	t, err := strconv.Atoi(s.String())
	v = uint32(t)
	return
}

func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}

func (s StrTo) Uint64() (res uint64, err error) {
	atoi, err := strconv.Atoi(s.String())
	res = uint64(atoi)
	return
}
