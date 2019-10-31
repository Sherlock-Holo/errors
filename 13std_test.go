// +build go1.13

package errors

import (
	stderrors "errors"
	"fmt"
	"testing"
)

func TestWithStack_Is_StdCompatibility(t *testing.T) {
	err := New("test")
	stackErr := WithStack(err)

	if !stderrors.Is(stackErr, err) {
		t.Error("not std compatible")
	}

	if !Is(stackErr, err) {
		t.Error("Is not works")
	}
}

func TestWithMessage_Is_StdCompatibility(t *testing.T) {
	err := New("test")
	msgErr := WithMessage(err, "message")

	if !stderrors.Is(msgErr, err) {
		t.Error("not std compatible")
	}

	if !Is(msgErr, err) {
		t.Error("Is not works")
	}
}

type customError struct{}

func (customError) Error() string {
	return ""
}

func TestWithStack_As_StdCompatibility(t *testing.T) {
	customErr := customError{}
	stackErr := WithStack(customErr)

	var ce customError
	if !stderrors.As(stackErr, &ce) {
		t.Error("not std compatible")
	}

	if !As(stackErr, &ce) {
		t.Error("As not works")
	}
}

func TestWithMessage_As_StdCompatibility(t *testing.T) {
	customErr := customError{}
	msgErr := WithMessage(customErr, "message")

	var ce customError
	if !stderrors.As(msgErr, &ce) {
		t.Error("not std compatible")
	}

	if !As(msgErr, &ce) {
		t.Error("As not works")
	}
}

func TestIs_StdCompatibility(t *testing.T) {
	err := stderrors.New("test")
	wrapErr := fmt.Errorf("wrap: %w", err)

	if !Is(wrapErr, err) {
		t.Error("Is not std compatible")
	}
}

func TestUnwrap_StdCompatibility(t *testing.T) {
	err := stderrors.New("test")
	wrapErr := fmt.Errorf("wrap: %w", err)

	if Unwrap(wrapErr) != err {
		t.Error("Is not std compatible")
	}
}

func TestAs_StdCompatibility(t *testing.T) {
	err := customError{}
	wrapErr := fmt.Errorf("wrap: %w", err)

	var ce customError
	if !As(wrapErr, &ce) {
		t.Error("As not std compatible")
	}
}
