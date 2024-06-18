package series

import (
	"fmt"
	"math"
	"strconv"
)

type int64Element struct {
	e   int64
	nan bool
}

// force uint64Element struct to implement Element interface
var _ Element = (*int64Element)(nil)

func (e *int64Element) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
		// i, err := strconv.Atoi(value.(string))
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			e.nan = true
			return
		}
		e.e = int64(i)
	case int:
		e.e = int64(val)
	case uint8:
		e.e = int64(val)
	case uint16:
		e.e = int64(val)
	case uint32:
		e.e = int64(val)
	case int64:
		e.e = int64(val)
	case uint64:
		e.e = int64(val)
	case float64:
		f := val
		if math.IsNaN(f) ||
			math.IsInf(f, 0) ||
			math.IsInf(f, 1) {
			e.nan = true
			return
		}
		e.e = int64(f)
	case float32:
		f := val
		if math.IsNaN(float64(f)) ||
			math.IsInf(float64(f), 0) ||
			math.IsInf(float64(f), 1) {
			e.nan = true
			return
		}
		e.e = int64(f)
	case bool:
		b := val
		if b {
			e.e = 1
		} else {
			e.e = 0
		}
	case Element:
		v, err := val.Int64()
		if err != nil {
			e.nan = true
			return
		}
		e.e = v
	default:
		e.nan = true
		return
	}
}

func (e int64Element) Copy() Element {
	if e.IsNA() {
		return &int64Element{0, true}
	}
	return &int64Element{e.e, false}
}

func (e int64Element) IsNA() bool {
	return e.nan
}

func (e int64Element) Type() Type {
	return Int
}

func (e int64Element) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return int(e.e)
}

func (e int64Element) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e int64Element) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(e.e), nil
}

func (e int64Element) Uint8() (uint8, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint8")
	}
	return uint8(e.e), nil
}

func (e int64Element) Uint16() (uint16, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint16")
	}
	return uint16(e.e), nil
}

func (e int64Element) Uint32() (uint32, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint32")
	}
	return uint32(e.e), nil
}

func (e int64Element) Int64() (int64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int64")
	}
	return e.e, nil
}

func (e int64Element) Uint64() (uint64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint64")
	}
	return uint64(e.e), nil
}

func (e int64Element) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return float64(e.e)
}

func (e int64Element) Float32() float32 {
	if e.IsNA() {
		return float32(math.NaN())
	}
	return float32(e.e)
}

func (e int64Element) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch e.e {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, fmt.Errorf("can't convert Uint64 \"%v\" to bool", e.e)
}

func (e int64Element) Eq(elem Element) bool {
	i, err := elem.Int64()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e == i
}

func (e int64Element) Neq(elem Element) bool {
	i, err := elem.Int64()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e != i
}

func (e int64Element) Less(elem Element) bool {
	i, err := elem.Int64()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e < i
}

func (e int64Element) LessEq(elem Element) bool {
	i, err := elem.Int64()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e <= i
}

func (e int64Element) Greater(elem Element) bool {
	i, err := elem.Int64()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e > i
}

func (e int64Element) GreaterEq(elem Element) bool {
	i, err := elem.Int64()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e >= i
}
