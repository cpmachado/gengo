package gengo

import (
	"errors"
	"math/rand"
	"time"
)

var (
	rnd                *rand.Rand
	ErrorNilFunc       = errors.New("f cannot be nil")
	ErrorFuncNilResult = errors.New("f returned nil")
)

var dtMin, dtMax int64

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	dtMin = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	dtMax = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC).Unix()
}

func Reseed() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func ReseedFunc(f func() *rand.Rand) error {
	if f == nil {
		return ErrorNilFunc
	}

	newRnd := f()
	if newRnd == nil {
		return ErrorFuncNilResult
	}
	rnd = newRnd

	return nil
}
