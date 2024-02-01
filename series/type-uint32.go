package series

import (
	"fmt"
	"math"
	"strconv"
)

type uint32Element struct {
	e   uint32
	nan bool
}

// force uint32Element struct to implement Element interface
var _ Element = (*uint32Element)(nil)

func (e *uint32Element) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
		// i, err := strconv.Atoi(value.(string))
		i, err := strconv.ParseUint(value.(string), 10, 32)
		if err != nil {
			e.nan = true
			return
		}
		e.e = uint32(i)
	case int:
		e.e = uint32(val)
	case uint8:
		e.e = uint32(val)
	case uint32:
		e.e = uint32(val)
	case uint64:
		e.e = uint32(val)
	case float64:
		f := val
		if math.IsNaN(f) ||
			math.IsInf(f, 0) ||
			math.IsInf(f, 1) {
			e.nan = true
			return
		}
		e.e = uint32(f)
	case float32:
		f := val
		if math.IsNaN(float64(f)) ||
			math.IsInf(float64(f), 0) ||
			math.IsInf(float64(f), 1) {
			e.nan = true
			return
		}
		e.e = uint32(f)
	case bool:
		b := val
		if b {
			e.e = 1
		} else {
			e.e = 0
		}
	case Element:
		v, err := val.Uint32()
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

func (e uint32Element) Copy() Element {
	if e.IsNA() {
		return &uint32Element{0, true}
	}
	return &uint32Element{e.e, false}
}

func (e uint32Element) IsNA() bool {
	return e.nan
}

func (e uint32Element) Type() Type {
	return Int
}

func (e uint32Element) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return int(e.e)
}

func (e uint32Element) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return fmt.Sprint(e.e)
}

func (e uint32Element) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(e.e), nil
}

func (e uint32Element) Uint8() (uint8, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint8")
	}
	return uint8(e.e), nil
}

func (e uint32Element) Uint32() (uint32, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint32")
	}
	return uint32(e.e), nil
}

func (e uint32Element) Uint64() (uint64, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to uint64")
	}
	return uint64(e.e), nil
}

func (e uint32Element) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return float64(e.e)
}

func (e uint32Element) Float32() float32 {
	if e.IsNA() {
		return float32(math.NaN())
	}
	return float32(e.e)
}

func (e uint32Element) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	switch e.e {
	case 1:
		return true, nil
	case 0:
		return false, nil
	}
	return false, fmt.Errorf("can't convert Uint32 \"%v\" to bool", e.e)
}

func (e uint32Element) Eq(elem Element) bool {
	i, err := elem.Uint32()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e == i
}

func (e uint32Element) Neq(elem Element) bool {
	i, err := elem.Uint32()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e != i
}

func (e uint32Element) Less(elem Element) bool {
	i, err := elem.Uint32()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e < i
}

func (e uint32Element) LessEq(elem Element) bool {
	i, err := elem.Uint32()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e <= i
}

func (e uint32Element) Greater(elem Element) bool {
	i, err := elem.Uint32()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e > i
}

func (e uint32Element) GreaterEq(elem Element) bool {
	i, err := elem.Uint32()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e >= i
}
