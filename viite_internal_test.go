package viite

import (
	"errors"
	"testing"
)

func TestRandomWithGenerateWithUnknownError(t *testing.T) {
	defer func() {
		randomGenerate = Generate
		_ = recover()
	}()
	randomGenerate = func(string) (string, error) {
		return "", errors.New("unknown error")
	}
	_ = Random()
	t.Errorf("expecting panic, but didn't happen")
}

func TestRandomWithGenerateWithErrInvalidInput(t *testing.T) {
	defer func() {
		randomGenerate = Generate
	}()
	expected := "123"
	randomGenerate = func(string) (string, error) {
		return expected, ErrInvalidInput
	}
	v := Random()
	if v != expected {
		t.Errorf("expecting '%s', got: '%v'", expected, v)
	}
}
