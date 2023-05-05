package number

import "strconv"

type Integer interface {
	~int8 | ~uint8 | ~int | ~uint | ~int32 | ~uint32 | ~int64 | ~uint64
}

func ParseInt[T Integer](str string) (T, error) {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return T(val), nil
}
