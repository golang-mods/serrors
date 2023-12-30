package serrors_test

import (
	"log/slog"

	"github.com/golang-mods/serrors"
)

func ExampleNew() {
	err := serrors.New("message", "foo", "bar")

	slog.Error(err.Error(), serrors.Attributes(err)...)
}

func ExampleFormat() {
	err := serrors.New("message", "foo", "bar")
	err = serrors.Format("%w: %s", err, "test")("baz", 1)

	slog.Error(err.Error(), serrors.Attributes(err)...)
}
